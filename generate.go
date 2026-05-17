package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	err := validate(input)
	if err != nil {
		return err.Error() // bad practice tho just for the purpose of recoding
	}
	lines := split(input)
	var results strings.Builder

	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				results.WriteString("/n")
			}
			continue
		}
		rows := RenderLine(line, banner)
		for _, row := range rows {
			results.WriteString(row + "\n")
		}
	}
	return results.String()
}
