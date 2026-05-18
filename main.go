package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run . \"text\"")
	}
	input := os.Args[1]
	text := strings.ReplaceAll(input, `n`, "\n")
	m, err := loadBanner("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := GenerateArt(text, m)
	fmt.Print(result)

}
