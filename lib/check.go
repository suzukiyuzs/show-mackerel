package showmackerel

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/mackerelio/mackerel-agent/config"
)

func parseCheckPlugin(data [][]string, file string) ([][]string, error) {
	load, err := config.LoadConfig(file)
	if err != nil {
		err := errors.New("Mackerel agent checkplugin file load error")
		return data, err
	}

	length := len(load.CheckPlugins)
	if length == 0 {
		err := errors.New("Mackerel agent checkplugin setting not found in file")
		return data, err
	}

	slice := []string{}
	for name, _ := range load.CheckPlugins {
		slice = append(slice, name)
	}
	sort.Strings(slice)

	for _, param := range slice {
		checkPlugin := load.CheckPlugins[param]
		cmd := checkPlugin.Command.Cmd
		notInterval := checkPlugin.NotificationInterval
		maxCheck := checkPlugin.MaxCheckAttempts
		chkInterval := checkPlugin.CheckInterval
		timeout := checkPlugin.Command.TimeoutDuration
		autoClose := checkPlugin.PreventAlertAutoClose
		env := checkPlugin.Command.Env
		action := checkPlugin.Action
		memo := checkPlugin.Memo

		var pCmd string
		pCmd = parseCheckPluginCommand(cmd)

		pNotInterval := "---"
		if notInterval != nil {
			pNotInterval = fmt.Sprint(*notInterval)
		}

		pMaxCheck := "---"
		if maxCheck != nil {
			pMaxCheck = fmt.Sprint(*maxCheck)
		}

		pChkInterval := "---"
		if chkInterval != nil {
			pChkInterval = fmt.Sprint(*chkInterval)
		}

		pTimeout := fmt.Sprint(timeout)
		pAutoClose := fmt.Sprint(autoClose)

		var pEnv string
		pEnv = parseSliceJoinComma(env)

		var pActionCmd, pActionTimeout, pActionEnv string
		if action != nil {
			actionCmd := checkPlugin.Action.Cmd
			actionTimeout := checkPlugin.Action.TimeoutDuration
			actionEnv := checkPlugin.Action.Env

			pActionCmd = parseCheckPluginCommand(actionCmd)
			pActionTimeout = fmt.Sprint(actionTimeout)
			pActionEnv = parseSliceJoinComma(actionEnv)
		}

		var pMemo string
		if memo != nil {
			mString := *memo
			mFields := strings.Fields(mString)
			pMemo = strings.Join(mFields, " ")
		} else {
			pMemo = ""
		}

		data = append(data, []string{
			param,
			pCmd,
			pNotInterval,
			pMaxCheck,
			pChkInterval,
			pTimeout,
			pAutoClose,
			pEnv,
			pActionCmd,
			pActionTimeout,
			pActionEnv,
			pMemo,
		})
	}

	return data, nil
}
