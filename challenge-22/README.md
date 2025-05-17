[View the Scoreboard](SCOREBOARD.md)

# Challenge 22: Greedy Coin Change

## Problem Statement

Implement a coin change algorithm that finds the minimum number of coins needed to make a given amount of change. You'll be using a greedy approach, which works with a specific set of coin denominations.

In this challenge, you'll work with the following coin denominations: `[1, 5, 10, 25, 50]` (representing penny, nickel, dime, quarter, and half-dollar coins).

You'll implement two functions:

1. `MinCoins` - Find the minimum number of coins needed to make change for a given amount.
2. `CoinCombination` - Find the specific combination of coins that gives the minimum number.

## Function Signatures

```go
func MinCoins(amount int, denominations []int) int
func CoinCombination(amount int, denominations []int) map[int]int
```

## Input Format

- `amount` - An integer representing the amount of change needed (in cents).
- `denominations` - A slice of integers representing the available coin denominations, sorted in ascending order.

## Output Format

- `MinCoins` - Returns an integer representing the minimum number of coins needed.
- `CoinCombination` - Returns a map where keys are coin denominations and values are the number of coins used for each denomination.

## Requirements

1. The `MinCoins` function should return the minimum number of coins needed to make the given amount.
2. The `CoinCombination` function should return a map with the specific combination of coins.
3. If the amount cannot be made with the given denominations, `MinCoins` should return -1 and `CoinCombination` should return an empty map.
4. Your solution should implement the greedy approach, which always chooses the largest coin possible.

## Sample Input and Output

### Sample Input 1

```
MinCoins(87, []int{1, 5, 10, 25, 50})
```

### Sample Output 1

```
6
```

### Sample Input 2

```
CoinCombination(87, []int{1, 5, 10, 25, 50})
```

### Sample Output 2

```
map[1:2 10:1 25:1 50:1]
```
(This represents 1 half-dollar, 1 quarter, 1 dime, and 2 pennies)

### Sample Input 3

```
MinCoins(42, []int{1, 5, 10, 25, 50})
```

### Sample Output 3

```
5
```

### Sample Input 4

```
CoinCombination(42, []int{1, 5, 10, 25, 50})
```

### Sample Output 4

```
map[1:2 5:1 10:1 25:1]
```
(This represents 1 quarter, 1 dime, 1 nickel, and 2 pennies)

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-22/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required functions.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-22/` directory:

```bash
go test -v
```

## Note About the Greedy Approach

The greedy approach for coin change works optimally for the standard U.S. coin denominations. However, it doesn't always produce the optimal result for all sets of denominations. For example, with denominations [1, 3, 4], the greedy approach would use 6 coins (4 + 1 + 1) to make 6, while the optimal solution is 2 coins (3 + 3). 