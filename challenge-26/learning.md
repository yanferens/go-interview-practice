# Learning Materials for Regular Expression Processor

## Regular Expressions in Go

This challenge focuses on using Go's `regexp` package to implement pattern matching and text manipulation.

### Understanding Regular Expressions in Go

Regular expressions (regex) are powerful patterns used to match character combinations in strings. Go provides robust regex support through the standard library's `regexp` package.

```go
import "regexp"
```

### Basic Pattern Matching

The simplest way to use regular expressions is to check if a pattern exists in a string:

```go
// Compile a regular expression pattern
pattern := `\d+`  // Match one or more digits
re, err := regexp.Compile(pattern)
if err != nil {
    // Handle error
}

// Check if the pattern matches a string
matched := re.MatchString("123 Main St")  // Returns true
```

For quick operations, you can use the package-level function:

```go
matched, err := regexp.MatchString(`\d+`, "123 Main St")
```

### Common Regex Patterns

Here are some frequently used patterns:

| Pattern | Description | Example |
|---------|-------------|---------|
| `\d` | Match a digit character | `\d+` matches "42" in "The answer is 42" |
| `\D` | Match a non-digit character | `\D+` matches "The answer is " in "The answer is 42" |
| `\w` | Match a word character (alphanumeric + underscore) | `\w+` matches "Hello_123" in "Hello_123!" |
| `\W` | Match a non-word character | `\W+` matches "!@#" in "abc!@#def" |
| `\s` | Match a whitespace character | `\s+` matches spaces between words |
| `\S` | Match a non-whitespace character | `\S+` matches each word |
| `.` | Match any character except newline | `a.b` matches "acb", "adb", etc. |
| `^` | Match the start of a string | `^Hello` matches "Hello" only at the beginning |
| `$` | Match the end of a string | `world$` matches "world" only at the end |
| `[abc]` | Match any character in the set | `[aeiou]` matches any vowel |
| `[^abc]` | Match any character not in the set | `[^aeiou]` matches any non-vowel |
| `a*` | Match 0 or more occurrences of 'a' | `a*` matches "", "a", "aa", etc. |
| `a+` | Match 1 or more occurrences of 'a' | `a+` matches "a", "aa", etc. but not "" |
| `a?` | Match 0 or 1 occurrence of 'a' | `a?` matches "" or "a" |
| `a{n}` | Match exactly n occurrences of 'a' | `a{3}` matches "aaa" |
| `a{n,}` | Match n or more occurrences of 'a' | `a{2,}` matches "aa", "aaa", etc. |
| `a{n,m}` | Match between n and m occurrences of 'a' | `a{2,4}` matches "aa", "aaa", or "aaaa" |
| `a\|b` | Match either 'a' or 'b' | `cat\|dog` matches "cat" or "dog" |
| `()` | Group patterns and capture matches | `(abc)+` matches "abc", "abcabc", etc. |

### Finding Matches

To find all matches in a string:

```go
re := regexp.MustCompile(`\d+`)  // MustCompile panics on error
matches := re.FindAllString("There are 15 apples and 25 oranges", -1)
// matches = ["15", "25"]
```

To find the first match:

```go
match := re.FindString("There are 15 apples and 25 oranges")
// match = "15"
```

### Capturing Groups

Capturing groups allow you to extract specific parts of a match:

```go
re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
matches := re.FindStringSubmatch("contact us at user@example.com for more info")
// matches = ["user@example.com", "user", "example", "com"]
```

The first element is the full match, followed by each captured group.

To get the indexes of matches:

```go
indexes := re.FindStringSubmatchIndex("contact us at user@example.com")
// indexes contains the start and end positions of each match
```

### Replacing Matches

Replace all matches with a new string:

```go
re := regexp.MustCompile(`\d+`)
result := re.ReplaceAllString("There are 15 apples and 25 oranges", "XX")
// result = "There are XX apples and XX oranges"
```

Use a function to dynamically determine replacements:

```go
re := regexp.MustCompile(`\d+`)
result := re.ReplaceAllStringFunc("15 apples and 25 oranges", func(s string) string {
    // Convert string number to int, double it, and return as string
    n, _ := strconv.Atoi(s)
    return strconv.Itoa(n * 2)
})
// result = "30 apples and 50 oranges"
```

### Using Named Capture Groups

Named capture groups make your regex more readable:

```go
re := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
matches := re.FindStringSubmatch("The date is 2023-11-15")
// matches = ["2023-11-15", "2023", "11", "15"]

// Extract by name
yearIndex := re.SubexpIndex("year")
year := matches[yearIndex]  // "2023"
```

### Splitting Strings

Split a string based on a regex pattern:

```go
re := regexp.MustCompile(`\s+`)  // Match one or more whitespace characters
parts := re.Split("Hello   world  !  ", -1)
// parts = ["Hello", "world", "!", ""]
```

### Compiling Flags

Modify regex behavior with flags:

```go
// Case-insensitive matching
re := regexp.MustCompile(`(?i)hello`)
matched := re.MatchString("Hello")  // true

// Multiline mode - ^ and $ match start/end of line
re := regexp.MustCompile(`(?m)^start`)
matched := re.MatchString("line1\nstart of line2")  // true

// Both flags
re := regexp.MustCompile(`(?im)^hello$`)
```

### Validating Input Formats

Regular expressions are excellent for validating input formats:

```go
// Email validation (simplified)
emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
valid := emailRegex.MatchString("user@example.com")  // true

// Phone number validation (US format)
phoneRegex := regexp.MustCompile(`^\(\d{3}\) \d{3}-\d{4}$`)
valid := phoneRegex.MatchString("(555) 123-4567")  // true

// URL validation (simplified)
urlRegex := regexp.MustCompile(`^https?://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9_.-]*)*$`)
valid := urlRegex.MatchString("https://www.example.com/path")  // true
```

### Performance Considerations

1. **Compile Once, Use Many Times**: Compiling a regex is expensive, so compile once and reuse the compiled pattern.

```go
// Bad practice
for _, s := range strings {
    matched, _ := regexp.MatchString(pattern, s)  // Recompiles every iteration
}

// Good practice
re := regexp.MustCompile(pattern)
for _, s := range strings {
    matched := re.MatchString(s)  // Reuses compiled pattern
}
```

2. **Avoid Excessive Backtracking**: Complex patterns with many repetition operators can cause excessive backtracking, leading to poor performance.

3. **Use Specific Patterns**: More specific patterns usually perform better than overly general ones.

### Working with Complex Patterns

For complex text processing, break down the task into multiple regexes:

```go
// Extract an HTML tag and its content
re := regexp.MustCompile(`<([a-z]+)>(.*?)</\1>`)
matches := re.FindAllStringSubmatch("<p>First paragraph</p><div>Content</div>", -1)
// matches = [["<p>First paragraph</p>", "p", "First paragraph"], ["<div>Content</div>", "div", "Content"]]
```

### Regular Expression Debugging

When debugging complex regexes, break them down and test each part:

```go
// Complex regex
fullPattern := `^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})([+-]\d{2}:\d{2})$`

// Break it down
datePattern := `^\d{4}-\d{2}-\d{2}`
timePattern := `\d{2}:\d{2}:\d{2}`
timezonePattern := `[+-]\d{2}:\d{2}$`

// Test each part
dateRe := regexp.MustCompile(datePattern)
timeRe := regexp.MustCompile(timePattern)
tzRe := regexp.MustCompile(timezonePattern)

dateMatches := dateRe.FindString("2023-11-15T14:30:45+02:00")  // "2023-11-15"
timeMatches := timeRe.FindString("2023-11-15T14:30:45+02:00")  // "14:30:45"
tzMatches := tzRe.FindString("2023-11-15T14:30:45+02:00")      // "+02:00"
```

### Practice Example: Log Analyzer

Here's a practical example that uses regex to parse log entries:

```go
package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {
    logLines := []string{
        "2023-11-15 14:23:45 INFO Server started on port 8080",
        "2023-11-15 14:24:12 ERROR Failed to connect to database: timeout",
        "2023-11-15 14:25:01 WARNING High memory usage: 85%",
    }

    // Define regex to parse log entries
    logPattern := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}) (\d{2}:\d{2}:\d{2}) (\w+) (.+)$`)

    for _, line := range logLines {
        matches := logPattern.FindStringSubmatch(line)
        if len(matches) == 5 {
            date := matches[1]
            time := matches[2]
            level := matches[3]
            message := matches[4]

            fmt.Printf("Date: %s, Time: %s, Level: %s, Message: %s\n", 
                      date, time, level, message)
            
            // Additional processing for error logs
            if level == "ERROR" {
                errorPattern := regexp.MustCompile(`Failed to (.*?): (.*)`)
                errorMatches := errorPattern.FindStringSubmatch(message)
                if len(errorMatches) == 3 {
                    action := errorMatches[1]
                    reason := errorMatches[2]
                    fmt.Printf("  Error details - Action: %s, Reason: %s\n", action, reason)
                }
            }
        }
    }
}
```

### Further Reading

- [Go regexp Package Documentation](https://golang.org/pkg/regexp/)
- [Regular Expressions Quick Start](https://github.com/google/re2/wiki/Syntax)
- [Regular Expression Testing Tool](https://regex101.com/)
- [The Go Programming Language - Chapter on Regular Expressions](https://www.gopl.io/) 