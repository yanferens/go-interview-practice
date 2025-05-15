# Learning Materials for Word Frequency Counter

## Maps in Go

Maps are a built-in data structure in Go that associate keys with values, similar to dictionaries, hash tables, or associative arrays in other languages.

### Basic Map Operations

```go
// Declare a map with string keys and integer values
wordFrequency := make(map[string]int)

// Set a value
wordFrequency["hello"] = 1

// Update a value
wordFrequency["hello"]++

// Get a value
count := wordFrequency["hello"]

// Check if a key exists
count, exists := wordFrequency["world"]
if exists {
    fmt.Println("'world' found with count:", count)
} else {
    fmt.Println("'world' not found in map")
}

// Delete a key
delete(wordFrequency, "hello")

// Iterate over a map
for word, count := range wordFrequency {
    fmt.Printf("%s: %d\n", word, count)
}
```

## String Handling in Go

Go provides several functions for string manipulation in the `strings` package.

### Common String Operations

```go
import "strings"

// Convert to lowercase
lowercase := strings.ToLower("Hello")  // "hello"

// Split a string
words := strings.Split("hello world", " ")  // ["hello", "world"]

// Join strings
joined := strings.Join([]string{"hello", "world"}, " ")  // "hello world"

// Replace all occurrences
replaced := strings.ReplaceAll("hello, hello!", ",", "")  // "hello hello!"

// Contains substring
hasPrefix := strings.Contains("hello world", "hello")  // true

// Trim whitespace
trimmed := strings.TrimSpace("  hello  ")  // "hello"
```

## Regular Expressions for Advanced String Processing

For more complex string processing, Go's `regexp` package provides powerful pattern matching:

```go
import "regexp"

// Create a regex to match only letters and digits
re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

// Replace all non-alphanumeric characters with a space
cleaned := re.ReplaceAllString("Hello, world! 123", " ")  // "Hello world 123"

// Split using regex
words := re.Split("Hello,world!123", -1)  // ["Hello", "world", "123"]
```

## Efficiency Considerations

When processing large texts:

1. **Pre-allocation**: If you know the approximate size of your map, initialize it with capacity:
   ```go
   wordFrequency := make(map[string]int, 1000)  // Pre-allocate for 1000 words
   ```

2. **Builder Pattern**: For complex string manipulation, use `strings.Builder`:
   ```go
   var builder strings.Builder
   for _, word := range words {
       builder.WriteString(word)
       builder.WriteString(" ")
   }
   result := builder.String()
   ```

3. **Single-pass Processing**: Avoid multiple iterations over the same data when possible

## Example: Improved Word Frequency Counter

```go
func CountWordFrequency(text string) map[string]int {
    // Convert to lowercase once
    text = strings.ToLower(text)
    
    // Use regex to replace all non-alphanumeric characters with spaces
    re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
    text = re.ReplaceAllString(text, " ")
    
    // Split by spaces and count
    words := strings.Fields(text)  // Fields splits by whitespace
    wordFrequency := make(map[string]int, len(words)/2)
    
    for _, word := range words {
        if word != "" {
            wordFrequency[word]++
        }
    }
    
    return wordFrequency
}
```

## Related Go Concepts

- **Hash Maps**: Go's maps are implemented as hash tables, providing O(1) average lookup time
- **Strings as UTF-8**: Go strings are UTF-8 encoded by default, so be careful when handling non-ASCII characters
- **Immutability**: Strings in Go are immutable, so operations like `ToLower()` create new strings
- **Runes**: For proper Unicode character handling, consider working with runes (`[]rune`) instead of bytes

## Further Reading

- [Go Maps in Action](https://blog.golang.org/maps)
- [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
- [Regular Expressions in Go](https://gobyexample.com/regular-expressions) 