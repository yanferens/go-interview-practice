# Challenge 26: Regular Expression Text Processor

## Problem Statement

In this challenge, you will implement a text processing utility that uses regular expressions to extract, validate, and transform data from various text formats.

Your task is to create a regular expression processor that can:

1. Extract specific data patterns from text (emails, phone numbers, dates, etc.)
2. Validate if input strings match specific formats
3. Replace or transform text based on pattern matching
4. Parse structured text like logs or CSV data

## Function Signatures

You will need to implement the following functions:

```go
// ExtractEmails extracts all valid email addresses from a text
func ExtractEmails(text string) []string

// ValidatePhone checks if a string is a valid phone number in format (XXX) XXX-XXXX
func ValidatePhone(phone string) bool

// MaskCreditCard replaces all but the last 4 digits of a credit card number with "X"
// Example: "1234-5678-9012-3456" -> "XXXX-XXXX-XXXX-3456"
func MaskCreditCard(cardNumber string) string

// ParseLogEntry parses a log entry with format:
// "YYYY-MM-DD HH:MM:SS LEVEL Message"
// Returns a map with keys: "date", "time", "level", "message"
func ParseLogEntry(logLine string) map[string]string

// ExtractURLs extracts all valid URLs from a text
func ExtractURLs(text string) []string
```

## Input/Output Examples

### ExtractEmails
- Input: `"Contact us at support@example.com or sales@company.co.uk for more info."`
- Output: `["support@example.com", "sales@company.co.uk"]`

### ValidatePhone
- Input: `"(555) 123-4567"`
- Output: `true`
- Input: `"555-123-4567"`
- Output: `false`

### MaskCreditCard
- Input: `"1234-5678-9012-3456"`
- Output: `"XXXX-XXXX-XXXX-3456"`
- Input: `"1234567890123456"`
- Output: `"XXXXXXXXXXXX3456"`

### ParseLogEntry
- Input: `"2023-11-15 14:23:45 INFO Server started on port 8080"`
- Output: 
```go
map[string]string{
    "date":    "2023-11-15",
    "time":    "14:23:45",
    "level":   "INFO",
    "message": "Server started on port 8080",
}
```

### ExtractURLs
- Input: `"Visit https://golang.org and http://example.com/page?q=123 for more information."`
- Output: `["https://golang.org", "http://example.com/page?q=123"]`

## Constraints

- Your solution should handle edge cases appropriately
- Regular expressions should be efficient and avoid excessive backtracking
- Compile regular expressions once and reuse them for better performance
- For email validation, use a reasonable regex that covers common email formats

## Evaluation Criteria

- Correctness: Does your solution handle all the required cases?
- Efficiency: Are your regular expressions optimized?
- Code Quality: Is your code well-structured and documented?
- Error Handling: Does your code handle invalid inputs gracefully?

## Learning Resources

See the [learning.md](learning.md) document for a comprehensive guide on using regular expressions in Go.

## Hints

1. The `regexp` package in Go provides comprehensive regular expression functionality
2. Use `MustCompile` for patterns you know are valid to simplify error handling
3. Remember to handle special characters in your patterns
4. For complex patterns, consider breaking them down into smaller parts
5. Test your regexes with a variety of inputs, including edge cases 