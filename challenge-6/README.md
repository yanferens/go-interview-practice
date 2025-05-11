[View the Scoreboard](SCOREBOARD.md)

# Challenge 6: Word Frequency Counter

## Problem Statement

Write a function `CountWordFrequency` that takes a string containing multiple words and returns a map where each key is a word and the value is the number of times that word appears in the string. The comparison should be case-insensitive, meaning "Hello" and "hello" should be counted as the same word.

## Function Signature

```go
func CountWordFrequency(text string) map[string]int
```

## Input Format

- A string `text` containing multiple words separated by spaces, punctuation, or line breaks.

## Output Format

- A map where keys are lowercase words and values are the frequency counts.

## Constraints

- Words are defined as sequences of letters and digits.
- All words should be converted to lowercase before counting.
- Ignore all punctuation, spaces, and other non-alphanumeric characters.
- `0 <= len(text) <= 10000`

## Sample Input and Output

### Sample Input 1

```
"The quick brown fox jumps over the lazy dog."
```

### Sample Output 1

```
{
    "the": 2,
    "quick": 1,
    "brown": 1,
    "fox": 1,
    "jumps": 1,
    "over": 1,
    "lazy": 1,
    "dog": 1
}
```

### Sample Input 2

```
"Hello, hello! How are you doing today? Today is a great day."
```

### Sample Output 2

```
{
    "hello": 2,
    "how": 1,
    "are": 1,
    "you": 1,
    "doing": 1,
    "today": 2,
    "is": 1,
    "a": 1,
    "great": 1,
    "day": 1
}
```

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-6/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the `CountWordFrequency` function.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-6/` directory:

```bash
go test -v
``` 