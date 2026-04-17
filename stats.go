package main

import "strings"

func countWords(text string) int {
	return len(strings.Fields(text))
}

func countCharacters(text string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range strings.ToLower(text) {
		m[r]++
	}
	return m
}
