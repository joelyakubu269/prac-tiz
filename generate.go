package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	var art strings.Builder
	lines := split(input)
	for i, line := range lines {
		if line == " " {

		}
	}
}
