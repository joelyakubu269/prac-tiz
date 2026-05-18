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
		return err.Error() // bad practice tho, we loose error type information with this, i used it just for recoding
	}
	lines := split(input)
	var result strings.Builder
	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 { // add a new line if we are not on the last line
				result.WriteString("\n")
				continue
			}
		}
		rows := RenderLine(line, banner) // rendering is done line by line
		for _, row := range rows {
			result.WriteString(row + "\n")
		}
	}
	return result.String()
}
