package showmackerel

import (
	"sort"
	"strings"
)

func parseCheckPluginCommand(c string) string {
	c = strings.Replace(c, "\\", "", -1)
	// Markdown
	c = strings.Replace(c, "|", "%PIPE%", -1)
	c = strings.Replace(c, "`", "%BSLASH%", -1)

	f := strings.Fields(c)
	p := strings.Join(f, " ")

	return p
}

func parseSliceJoinComma(s []string) string {
	p := ""
	if s != nil {
		sort.Strings(s)
		p = strings.Join(s, ", ")
	}

	return p
}
