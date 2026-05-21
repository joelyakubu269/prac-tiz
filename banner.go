package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(filename string) (map[rune][]string, error) {
	m := make(map[rune][]string)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading banner")
	}
	charlines := strings.Split(string(data), "\n")
	for i := 32; i <= 126; i++ {
		index := i - 32
		start := index * 9
		if start+8 > len(charlines) {
			return nil, fmt.Errorf("invalid banner format")
		}
		m[rune(i)] = charlines[start : start+8]
	}
	return m, nil
}
