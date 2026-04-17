package main

import (
	"cmp"
	"slices"
	"strings"
	"unicode"
)

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

type CharacterCount struct {
	Character rune
	Count     int
}

func sortedCharacterCounts(counts map[rune]int) []CharacterCount {
	characterCounts := make([]CharacterCount, 0, len(counts))
	for k, v := range counts {
		// filter out non-alphabetic characters
		if unicode.IsLetter(k) {
			characterCounts = append(characterCounts, CharacterCount{Character: k, Count: v})
		}
	}

	slices.SortFunc(characterCounts, func(a, b CharacterCount) int {
		// sort in descending order of count
		return cmp.Compare(b.Count, a.Count)
	})

	return characterCounts
}
