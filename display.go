package main

import (
	"fmt"
	"strings"
)

func FormatBytes(b []byte) (out string) {
	lines := []string{"[]byte{"}
	line := ""
	for i := 0; i < len(b); i++ {
		if len(line) > 40 {
			lines = append(lines, line)
			line = ""
		}
		line += fmt.Sprintf("%d,", b[i])
	}
	if line != "" {
		lines = append(lines, line)
	}
	lines = append(lines, "}")
	out = strings.Join(lines, "\n")
	return
}
