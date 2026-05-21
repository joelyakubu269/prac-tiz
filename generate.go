package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	var result string
	var art strings.Builder
	lines := split(input)
	for i, line := range lines {
		if line == " " { // this is for new line character as it was not considered in RenderLine function
			art.WriteString("\n")
			continue
		}
		rows := RenderLine(line, banner)
		for _, row := range rows {
			art.WriteString(row + "\n")
		}
		result[i]
	}
}
