package showmackerel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type MackerelMonitors struct {
	Monitors []struct {
		ID                   string   `json:"id"`
		Name                 string   `json:"name"`
		Memo                 string   `json:"memo"`
		Type                 string   `json:"type"`
		IsMute               bool     `json:"isMute"`
		NotificationInterval uint64   `json:"notificationInterval"`
		Metric               string   `json:"metric"`
		Operator             string   `json:"operator"`
		Warning              float64  `json:"warning"`
		Critical             float64  `json:"critical"`
		Duration             uint64   `json:"duration"`
		MaxCheckAttempts     uint64   `json:"maxCheckAttempts"`
		Scopes               []string `json:"scopes"`
		ExcludeScopes        []string `json:"excludeScopes"`
	} `json:"monitors"`
}

func parseMonitor(data [][]string, file string) ([][]string, error) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		err := errors.New("Mackerel monitor file load error")
		return data, err
	}

	var mon MackerelMonitors

	if err = json.Unmarshal(raw, &mon); err != nil {
		err := errors.New("Mackerel monitor file parse error")
		return data, err
	}

	for _, m := range mon.Monitors {
		pID := fmt.Sprint(m.ID)
		pName := fmt.Sprint(m.Name)

		pMemo := "---"
		if m.Memo != "" {
			pMemo = fmt.Sprint(m.Memo)
		}

		pType := fmt.Sprint(m.Type)
		pIsMute := fmt.Sprint(m.IsMute)
		pNotInterval := fmt.Sprint(m.NotificationInterval)

		pMetric := "---"
		if m.Metric != "" {
			pMetric = fmt.Sprint(m.Metric)
		}

		pOperator := "---"
		if m.Operator != "" {
			pOperator = fmt.Sprint(m.Operator)
		}

		pWarning := fmt.Sprint(m.Warning)
		pCritical := fmt.Sprint(m.Critical)
		pDuration := fmt.Sprint(m.Duration)
		pMaxChk := fmt.Sprint(m.MaxCheckAttempts)

		pScopes := "---"
		if len(m.Scopes) != 0 {
			pScopes = parseSliceJoinComma(m.Scopes)
		}

		pExScopes := "---"
		if len(m.ExcludeScopes) != 0 {
			pExScopes = parseSliceJoinComma(m.ExcludeScopes)
		}

		data = append(data, []string{
			pID,
			pName,
			pMemo,
			pType,
			pIsMute,
			pNotInterval,
			pMetric,
			pOperator,
			pWarning,
			pCritical,
			pDuration,
			pMaxChk,
			pScopes,
			pExScopes,
		})

	}

	return data, nil
}
