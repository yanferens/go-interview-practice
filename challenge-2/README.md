# Scoreboard for challenge-2
| Username   | Passed Tests | Total Tests |
|------------|--------------|-------------|
| RezaSi | 7 | 7 |


[View the Scoreboard](SCOREBOARD.md)

# Challenge 2: Reverse a String

## Problem Statement

Write a function `ReverseString` that takes a string and returns the string reversed.

## Function Signature

```go
func ReverseString(s string) string
```

## Input Format

- A single line containing a string `s`.

## Output Format

- The reversed string.

## Constraints

- `0 <= len(s) <= 1000`
- The string may contain ASCII letters, digits, and special characters.

## Sample Input and Output

### Sample Input 1

```
hello
```

### Sample Output 1

```
olleh
```

### Sample Input 2

```
Go is fun!
```

### Sample Output 2

```
!nuf si oG
```

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-2/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the `ReverseString` function.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

1. **Initialize Go Module** (if not already initialized):

   Navigate to the `challenge-2` directory:

   ```bash
   cd challenge-2
   go mod init challenge2
   ```

2. **Run the Tests:**

   ```bash
   go test -v
   ```
