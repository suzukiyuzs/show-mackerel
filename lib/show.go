package showmackerel

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/olekukonko/tablewriter"
)

var cmdOptions struct {
	Mackerel string `short:"m" long:"mackerel" required:"true" description:"Mackerel type(check or monitor)"`
	File     string `short:"f" long:"file" required:"true" description:"Mackerel file path"`
}

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func Show(args []string) int {
	opts := cmdOptions
	_, err := flags.ParseArgs(&opts, args[1:])
	if err != nil {
		return ExitCodeError
	}

	mackerel := opts.Mackerel
	file := opts.File

	data := [][]string{}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetColWidth(1000)
	table.SetAutoFormatHeaders(false)

	var title string
	var header []string
	if mackerel == "check" {
		title = "# Mackerel check plugin"
		header = []string{
			"Name",
			"Command",
			"NotificationInterval",
			"MaxCheckAttempts",
			"CheckInterval",
			"TimeoutSeconds",
			"PreventAlertAutoClose",
			"Env",
			"Action.Command",
			"Action.TimeoutSeconds",
			"Action.Env",
			"Memo",
		}
		data, err = parseCheckPlugin(data, file)
		if err != nil {
			fmt.Println(err)
			return ExitCodeError
		}

	} else if mackerel == "monitor" {
		title = "# Mackerel monitor"
		header = []string{
			"ID",
			"Name",
			"Memo",
			"Type",
			"IsMute",
			"NotificationInterval",
			"Metric",
			"Operator",
			"Warning",
			"Critical",
			"Duratoin",
			"MaxCheckAttempts",
			"Scopes",
			"ExcludeScopes",
		}
		data, err = parseMonitor(data, file)
		if err != nil {
			fmt.Println(err)
			return ExitCodeError
		}

	} else {
		fmt.Println("Mackerel type is miss match")
		return ExitCodeError
	}

	for _, d := range data {
		table.Append(d)
	}

	table.SetHeader(header)
	fmt.Println(title)
	table.Render()

	return ExitCodeOK
}
