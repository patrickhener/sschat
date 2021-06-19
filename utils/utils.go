package utils

import (
	"strings"

	markdown "github.com/quackduck/go-term-markdown"
)

// MDRender will utilise 'markdown' to render print style
func MDRender(msg string, before int, lineWidth int) string {
	md := string(markdown.Render(msg, lineWidth-(before), 0))
	md = strings.TrimSuffix(md, "\n")
	split := strings.Split(md, "\n")
	for i := range split {
		if i == 0 {
			continue // first line will be padded
		}
		split[i] = strings.Repeat(" ", before) + split[i]
	}
	if len(split) == 1 {
		return md
	}
	return strings.Join(split, "\n")
}
