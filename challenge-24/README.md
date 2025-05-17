[View the Scoreboard](SCOREBOARD.md)

# Challenge 24: Dynamic Programming - Longest Increasing Subsequence

## Problem Statement

The Longest Increasing Subsequence (LIS) problem is a classic dynamic programming problem. Given a sequence of integers, find the length of the longest subsequence such that all elements of the subsequence are sorted in increasing order. A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

In this challenge, you will implement three different approaches to solve the LIS problem:

1. `DPLongestIncreasingSubsequence` - A standard dynamic programming solution with O(n²) time complexity.
2. `OptimizedLIS` - An optimized solution with O(n log n) time complexity using binary search.
3. `GetLISElements` - A function that returns the actual elements of the LIS, not just its length.

## Function Signatures

```go
func DPLongestIncreasingSubsequence(nums []int) int
func OptimizedLIS(nums []int) int
func GetLISElements(nums []int) []int
```

## Input Format

- `nums` - A slice of integers representing the sequence.

## Output Format

- `DPLongestIncreasingSubsequence` - Returns an integer representing the length of the LIS.
- `OptimizedLIS` - Returns an integer representing the length of the LIS.
- `GetLISElements` - Returns a slice of integers representing the elements of one possible LIS.

## Requirements

1. `DPLongestIncreasingSubsequence` should implement the standard dynamic programming solution with O(n²) time complexity.
2. `OptimizedLIS` should implement an optimized solution with O(n log n) time complexity.
3. `GetLISElements` should return the actual elements of the LIS, not just its length.
4. Handle edge cases such as empty slices or slices with a single element.
5. If multiple LIS exist with the same length, returning any valid LIS is acceptable.

## Sample Input and Output

### Sample Input 1

```
DPLongestIncreasingSubsequence([]int{10, 9, 2, 5, 3, 7, 101, 18})
```

### Sample Output 1

```
4
```

### Sample Input 2

```
OptimizedLIS([]int{0, 1, 0, 3, 2, 3})
```

### Sample Output 2

```
4
```

### Sample Input 3

```
GetLISElements([]int{10, 9, 2, 5, 3, 7, 101, 18})
```

### Sample Output 3

```
[2, 5, 7, 101]
```
Note: [2, 3, 7, 18] or [2, 3, 7, 101] would also be valid outputs as they are also valid LIS.

### Sample Input 4

```
DPLongestIncreasingSubsequence([]int{7, 7, 7, 7, 7, 7, 7})
```

### Sample Output 4

```
1
```

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-24/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required functions.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-24/` directory:

```bash
go test -v
```

## Performance Expectations

- **DPLongestIncreasingSubsequence**: O(n²) time complexity, O(n) space complexity.
- **OptimizedLIS**: O(n log n) time complexity, O(n) space complexity.
- **GetLISElements**: O(n²) or O(n log n) time complexity, depending on the approach. 