package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("============ BOOKBOT ============")

	filename := "books/frankenstein.txt"
	fmt.Printf("Analyzing book found at %s...\n", filename)

	contents, err := getBookText(filename)
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	fmt.Println("----------- Word Count ----------")
	fmt.Println("Found", countWords(contents), "total words")

	fmt.Println("--------- Character Count -------")
	sortedCounts := sortedCharacterCounts(countCharacters(contents))
	for _, cc := range sortedCounts {
		fmt.Printf("%c: %d\n", cc.Character, cc.Count)
	}

	fmt.Println("============= END ===============")
}

func getBookText(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
