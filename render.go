package main

import (
	"strings"
)

func RenderLine(text string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i < 8; i++ {
		var rows strings.Builder
		for _, c := range text {
			if val, ok := banner[c]; ok {
				rows.WriteString(val[i])
			}
		}
		result[i] = rows.String
	}
}
