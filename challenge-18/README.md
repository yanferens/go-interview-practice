[View the Scoreboard](SCOREBOARD.md)

# Challenge 18: Temperature Converter

## Problem Statement

Write a program that converts temperatures between Celsius and Fahrenheit. You'll implement two functions:
1. `CelsiusToFahrenheit` - converts a temperature from Celsius to Fahrenheit.
2. `FahrenheitToCelsius` - converts a temperature from Fahrenheit to Celsius.

## Function Signatures

```go
func CelsiusToFahrenheit(celsius float64) float64
func FahrenheitToCelsius(fahrenheit float64) float64
```

## Input Format

- A float64 temperature value in either Celsius or Fahrenheit.

## Output Format

- A float64 temperature value converted to the other unit.

## Conversion Formulas

- **Celsius to Fahrenheit**: F = C × 9/5 + 32
- **Fahrenheit to Celsius**: C = (F - 32) × 5/9

## Sample Input and Output

### Sample Input 1

```
CelsiusToFahrenheit(0)
```

### Sample Output 1

```
32.0
```

### Sample Input 2

```
FahrenheitToCelsius(32)
```

### Sample Output 2

```
0.0
```

### Sample Input 3

```
CelsiusToFahrenheit(100)
```

### Sample Output 3

```
212.0
```

## Requirements

1. Round the result to 2 decimal places
2. Handle negative temperatures correctly
3. The functions should work with any valid temperature value

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-18/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required functions.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-18/` directory:

```bash
go test -v
``` 