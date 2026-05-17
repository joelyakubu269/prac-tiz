package main

import (
	"fmt"
)

func validate(text string) error {
	for _, r := range text {
		if r < 32 || r > 126 {
			return fmt.Errorf("invalid character")
		}
	}
	return nil
}
