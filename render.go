package main

import (
	"strings"
)

func RenderLine(text string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i < 8; i++ {
		var rows strings.Builder // creates a new empty builder every iteration, so that rows dont overlap
		for _, c := range text {
			if val, ok := banner[c]; ok {
				rows.WriteString(val[i])
			}
		}
		result[i] = rows.String() // collects the characters row by row
	}
	return result
}
