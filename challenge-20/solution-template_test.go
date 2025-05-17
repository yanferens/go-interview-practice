package main

import (
	"testing"
)

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Empty string", "", 0},
		{"Single word", "Go", 2},
		{"Multiple words", "Go is awesome", 12},
		{"With newlines", "Go\nis\nawesome", 12},
		{"With punctuation", "Hello, World!", 13},
		{"Complex text", "The Go programming language is an open source project to make programmers more productive.", 90},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountCharacters(tt.text)
			if result != tt.expected {
				t.Errorf("CountCharacters(%q) = %d, expected %d", tt.text, result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Empty string", "", 0},
		{"Single word", "Go", 1},
		{"Multiple words", "Go is awesome", 3},
		{"With newlines", "Go\nis\nawesome", 3},
		{"With punctuation", "Hello, World!", 2},
		{"Multiple spaces", "This   has   multiple   spaces", 4},
		{"Complex text", "The Go programming language is an open source project to make programmers more productive.", 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.text)
			if result != tt.expected {
				t.Errorf("CountWords(%q) = %d, expected %d", tt.text, result, tt.expected)
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Empty string", "", 1},
		{"Single line", "Go is awesome", 1},
		{"Two lines", "Go is awesome\nGo is cool", 2},
		{"Three lines", "Go\nis\nawesome", 3},
		{"Trailing newline", "Go is awesome\n", 2},
		{"Multiple newlines", "Go\n\nis\n\nawesome", 5},
		{"Complex text", "The Go programming language is an open source project to make programmers more productive.\n\nGo is expressive, concise, clean, and efficient.", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountLines(tt.text)
			if result != tt.expected {
				t.Errorf("CountLines(%q) = %d, expected %d", tt.text, result, tt.expected)
			}
		})
	}
}

func TestCountWordOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		word     string
		expected int
	}{
		{"Empty string", "", "go", 0},
		{"Word not found", "Hello, World!", "go", 0},
		{"Single occurrence", "Go is awesome", "go", 1},
		{"Multiple occurrences", "Go is cool. I love Go. Let's Go!", "go", 3},
		{"Case insensitive", "GO is not the same as go or Go or gO", "go", 4},
		{"Word boundaries", "Golang is not the same as Go language", "go", 1},
		{"Complex text with multiple occurrences", "The Go programming language is an open source project. Go is expressive and Go is efficient.", "go", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWordOccurrences(tt.text, tt.word)
			if result != tt.expected {
				t.Errorf("CountWordOccurrences(%q, %q) = %d, expected %d", tt.text, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWithExampleText(t *testing.T) {
	// First example from the challenge description
	example1 := `Go is an open source programming language that makes it easy to build
simple, reliable, and efficient software.`

	if count := CountCharacters(example1); count != 107 {
		t.Errorf("CountCharacters on example 1 = %d, expected 107", count)
	}

	if count := CountWords(example1); count != 17 {
		t.Errorf("CountWords on example 1 = %d, expected 17", count)
	}

	if count := CountLines(example1); count != 2 {
		t.Errorf("CountLines on example 1 = %d, expected 2", count)
	}

	if count := CountWordOccurrences(example1, "go"); count != 1 {
		t.Errorf("CountWordOccurrences(example1, \"go\") = %d, expected 1", count)
	}

	// Second example from the challenge description
	example2 := `The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to
write programs that get the most out of multicore and networked machines, while its novel type
system enables flexible and modular program construction.`

	if count := CountCharacters(example2); count != 313 {
		t.Errorf("CountCharacters on example 2 = %d, expected 313", count)
	}

	if count := CountWords(example2); count != 52 {
		t.Errorf("CountWords on example 2 = %d, expected 52", count)
	}

	if count := CountLines(example2); count != 4 {
		t.Errorf("CountLines on example 2 = %d, expected 4", count)
	}

	if count := CountWordOccurrences(example2, "go"); count != 3 {
		t.Errorf("CountWordOccurrences(example2, \"go\") = %d, expected 3", count)
	}
}

func TestSampleFile(t *testing.T) {
	// Create a sample file in memory
	sampleFile := `This is a sample text file.
It has multiple lines.
We can count characters, words, and lines.
Go is a great programming language to work with text.
Go makes string processing easy and efficient.
`

	// Test all functions with the sample file
	t.Run("Sample file tests", func(t *testing.T) {
		if count := CountCharacters(sampleFile); count != 187 {
			t.Errorf("CountCharacters = %d, expected 187", count)
		}

		if count := CountWords(sampleFile); count != 34 {
			t.Errorf("CountWords = %d, expected 34", count)
		}

		if count := CountLines(sampleFile); count != 6 {
			t.Errorf("CountLines = %d, expected 6", count)
		}

		if count := CountWordOccurrences(sampleFile, "go"); count != 2 {
			t.Errorf("CountWordOccurrences(\"go\") = %d, expected 2", count)
		}

		if count := CountWordOccurrences(sampleFile, "line"); count != 2 {
			t.Errorf("CountWordOccurrences(\"line\") = %d, expected 2", count)
		}
	})
}
