package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func loadWordsFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	// Read the file content
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := stat.Size()
	content := make([]byte, size)
	_, err = file.Read(content)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the file content into a string
	text := string(content)

	// Split the text into individual words
	words := strings.Fields(text)

	// Print the words
	return words
}

func main() {
	// Create a new Trie tree
	trie := InitTrie()

	// Load words
	for _, word := range loadWordsFromFile("words.txt") {
		trie.Insert(word)
	}

	// Print tree
	trie.PrintTrie()

	fmt.Println(fmt.Sprintf("Number of words: %d\n", trie.CountWords()))
}
