# Learning Materials for File Counter

## Working with Files and Text in Go

This challenge explores text processing and file handling in Go, focusing on basic string operations and counting techniques.

### Reading Files

Go provides several ways to read files:

#### Reading an Entire File at Once

```go
import (
    "io/ioutil"
    "fmt"
)

func readFile(filename string) {
    // Read the entire file at once
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    // Convert to string and print
    text := string(content)
    fmt.Println(text)
}
```

#### Reading a File Line by Line

```go
import (
    "bufio"
    "fmt"
    "os"
)

func readFileByLine(filename string) {
    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Create a scanner to read line by line
    scanner := bufio.NewScanner(file)
    lineCount := 0
    
    // Read line by line
    for scanner.Scan() {
        line := scanner.Text()
        lineCount++
        fmt.Printf("Line %d: %s\n", lineCount, line)
    }
    
    // Check for scanner errors
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}
```

### String Manipulation

Go provides rich support for working with strings through the `strings` package.

#### Counting Characters

To count the total number of characters in a string:

```go
func countCharacters(text string) int {
    return len(text)
}
```

#### Counting Words

To count words in a string, you can split the string by whitespace:

```go
import "strings"

func countWords(text string) int {
    if len(text) == 0 {
        return 0
    }
    
    // Split by whitespace and count non-empty elements
    words := strings.Fields(text)
    return len(words)
}
```

#### Counting Lines

To count lines in a string, you can count the newline characters and add 1:

```go
import "strings"

func countLines(text string) int {
    if len(text) == 0 {
        return 1
    }
    
    // Count newlines and add 1 (for the last line which might not end with a newline)
    return 1 + strings.Count(text, "\n")
}
```

#### Counting Word Occurrences

To count how many times a specific word appears in a text (case-insensitive):

```go
import (
    "strings"
    "regexp"
)

func countWordOccurrences(text, word string) int {
    if len(text) == 0 || len(word) == 0 {
        return 0
    }
    
    // Convert to lowercase for case-insensitive matching
    lowerText := strings.ToLower(text)
    lowerWord := strings.ToLower(word)
    
    // Method 1: Simple approach using strings.Count (but doesn't handle word boundaries)
    // return strings.Count(lowerText, lowerWord)
    
    // Method 2: Using regular expressions to match whole words only
    re := regexp.MustCompile(`\b` + regexp.QuoteMeta(lowerWord) + `\b`)
    return len(re.FindAllString(lowerText, -1))
}
```

### The `strings` Package

The `strings` package provides many useful functions for working with strings:

```go
import "strings"

// Check if a string contains a substring
contains := strings.Contains("Go is awesome", "awesome")  // true

// Split a string by a separator
parts := strings.Split("a,b,c", ",")  // ["a", "b", "c"]

// Join string slices with a separator
joined := strings.Join([]string{"a", "b", "c"}, "-")  // "a-b-c"

// Convert case
lower := strings.ToLower("Go")  // "go"
upper := strings.ToUpper("Go")  // "GO"

// Trim whitespace
trimmed := strings.TrimSpace(" Go ")  // "Go"

// Replace occurrences
replaced := strings.Replace("Go Go Go", "Go", "Golang", 2)  // "Golang Golang Go"
replacedAll := strings.ReplaceAll("Go Go Go", "Go", "Golang")  // "Golang Golang Golang"

// Count occurrences
count := strings.Count("Go Go Go", "Go")  // 3
```

### The `regexp` Package

For more complex string matching and manipulation, Go provides the `regexp` package:

```go
import (
    "fmt"
    "regexp"
)

func regexpExample() {
    text := "The Go programming language was announced in 2009."
    
    // Find all occurrences of "Go" as a whole word
    re := regexp.MustCompile(`\bGo\b`)
    matches := re.FindAllString(text, -1)
    fmt.Println(matches)  // ["Go"]
    
    // Find all words
    wordRe := regexp.MustCompile(`\b\w+\b`)
    words := wordRe.FindAllString(text, -1)
    fmt.Println(words)  // ["The", "Go", "programming", "language", "was", "announced", "in", "2009"]
    
    // Replace with regex
    replaced := re.ReplaceAllString(text, "Golang")
    fmt.Println(replaced)  // "The Golang programming language was announced in 2009."
}
```

### Command-Line Arguments

To read command-line arguments, you can use the `os.Args` slice:

```go
import (
    "fmt"
    "os"
)

func main() {
    // os.Args[0] is the program name
    // os.Args[1] is the first argument, etc.
    
    if len(os.Args) < 2 {
        fmt.Println("Please provide a filename as an argument")
        os.Exit(1)
    }
    
    filename := os.Args[1]
    fmt.Printf("Reading file: %s\n", filename)
}
```

## Handling Text Processing Efficiently

### Scanner vs. Split

When processing large files, using a `bufio.Scanner` is more memory-efficient than reading the entire file and then splitting it:

```go
import (
    "bufio"
    "os"
)

// Efficient line-by-line processing
func countLinesEfficiently(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineCount := 0
    
    for scanner.Scan() {
        lineCount++
    }
    
    return lineCount, scanner.Err()
}
```

## Further Reading

- [Go by Example: Reading Files](https://gobyexample.com/reading-files)
- [strings package documentation](https://pkg.go.dev/strings)
- [regexp package documentation](https://pkg.go.dev/regexp)
- [bufio package documentation](https://pkg.go.dev/bufio)
- [ioutil package documentation](https://pkg.go.dev/io/ioutil) 