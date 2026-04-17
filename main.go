package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	contents, err := getBookText("books/frankenstein.txt")
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	fmt.Println("Found", countWords(contents), "total words")

	count := countCharacters(contents)
	for k, v := range count {
		fmt.Printf("%c: %d\n", k, v)
	}
}

func getBookText(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
