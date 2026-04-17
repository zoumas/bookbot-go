package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	contents, err := getBookText("books/frankenstein.txt")
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	fmt.Println("Found", countWords(contents), "total words")
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

func getBookText(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
