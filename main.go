package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	contents, err := getBookText("books/frankenstein.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(contents)
}

func getBookText(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(contents), nil
}
