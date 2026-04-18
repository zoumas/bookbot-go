package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <book_file>")
		os.Exit(1)
	}

	fmt.Println("============ BOOKBOT ============")

	filename := os.Args[1]
	fmt.Printf("Analyzing book found at %s...\n", filename)

	contents, err := getBookText(filename)
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	fmt.Println("----------- Word Count ----------")
	fmt.Println("Found", countWords(contents), "total words")

	fmt.Println("--------- Character Count -------")
	sortedCounts := sortedCharacterCounts(countCharacters(contents))
	for _, c := range sortedCounts {
		fmt.Printf("%c: %d\n", c.Character, c.Count)
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
