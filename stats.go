package main

import "strings"

func countWords(text string) int {
	return len(strings.Fields(text))
}
