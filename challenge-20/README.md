[View the Scoreboard](SCOREBOARD.md)

# Challenge 20: File Counter

## Problem Statement

Write a program that reads a text file and counts the occurrences of various statistics:

1. The total number of characters
2. The total number of words
3. The total number of lines
4. The occurrences of a specific word (case-insensitive)

You'll implement the following functions:

## Function Signatures

```go
func CountCharacters(text string) int
func CountWords(text string) int
func CountLines(text string) int
func CountWordOccurrences(text string, word string) int
```

## Input Format

- For all functions, a string `text` containing the file contents.
- For `CountWordOccurrences`, an additional string `word` to count.

## Output Format

- For all functions, an integer representing the count.

## Requirements

1. `CountCharacters` should count all characters in the text, including whitespace and punctuation.
2. `CountWords` should count all words separated by whitespace.
3. `CountLines` should count all lines in the text (separated by newline characters).
4. `CountWordOccurrences` should count how many times a specific word appears in the text, ignoring case.

## Sample Input and Output

### Sample Input 1

```
Go is an open source programming language that makes it easy to build
simple, reliable, and efficient software.
```

### Sample Output 1

```
CountCharacters: 107
CountWords: 17
CountLines: 2
CountWordOccurrences("go"): 1
```

### Sample Input 2

```
The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to
write programs that get the most out of multicore and networked machines, while its novel type
system enables flexible and modular program construction.
```

### Sample Output 2

```
CountCharacters: 313
CountWords: 52
CountLines: 4
CountWordOccurrences("go"): 3
```

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-20/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required functions.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-20/` directory:

```bash
go test -v
``` 