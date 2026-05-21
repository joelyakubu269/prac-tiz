package main

import (
	"strings"
)

func RenderLine(text string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i < 8; i++ {
		var rows strings.Builder
		for _, r := range text {
			val, ok := banner[r]
			if ok {
				rows.WriteString(val[i]) // i is necessary in order to write values just perculiar to that row
			}
		}
		result[i] = rows.String()
	}
	return result
}
