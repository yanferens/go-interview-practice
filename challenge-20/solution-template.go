package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Check if a filename was provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		os.Exit(1)
	}

	// Get the filename from command line arguments
	filename := os.Args[1]

	// Read the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert content to string
	text := string(content)

	// Count statistics
	charCount := CountCharacters(text)
	wordCount := CountWords(text)
	lineCount := CountLines(text)

	// Print the results
	fmt.Printf("File: %s\n", filename)
	fmt.Printf("Character count: %d\n", charCount)
	fmt.Printf("Word count: %d\n", wordCount)
	fmt.Printf("Line count: %d\n", lineCount)

	// If a word to count was provided, count its occurrences
	if len(os.Args) >= 3 {
		wordToCount := os.Args[2]
		occurrences := CountWordOccurrences(text, wordToCount)
		fmt.Printf("Occurrences of '%s': %d\n", wordToCount, occurrences)
	}
}

// CountCharacters counts the total number of characters in the text
func CountCharacters(text string) int {
	// TODO: Implement this function
	return 0
}

// CountWords counts the total number of words in the text
func CountWords(text string) int {
	// TODO: Implement this function
	return 0
}

// CountLines counts the total number of lines in the text
func CountLines(text string) int {
	// TODO: Implement this function
	return 0
}

// CountWordOccurrences counts how many times a specific word appears in the text (case-insensitive)
func CountWordOccurrences(text string, word string) int {
	// TODO: Implement this function
	return 0
}
