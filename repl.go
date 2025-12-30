package main

import (
	"strings"
)

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lowerCase := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowerCase)
	textSlice := strings.Split(trimmed, " ")

	return textSlice
}
