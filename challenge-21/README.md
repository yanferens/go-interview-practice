[View the Scoreboard](SCOREBOARD.md)

# Challenge 21: Binary Search Implementation

## Problem Statement

Implement the binary search algorithm to efficiently find items in a sorted collection. Binary search is a divide-and-conquer algorithm that repeatedly divides the search space in half, making it much faster than linear search for sorted data.

You'll implement three versions of binary search:

1. `BinarySearch` - Standard binary search that returns the index of a target value.
2. `BinarySearchRecursive` - A recursive implementation of binary search.
3. `FindInsertPosition` - Find the position where a value should be inserted to maintain sorted order.

## Function Signatures

```go
func BinarySearch(arr []int, target int) int
func BinarySearchRecursive(arr []int, target int, left int, right int) int
func FindInsertPosition(arr []int, target int) int
```

## Input Format

- For all functions, a sorted slice of integers `arr` and a target integer value.
- For the recursive function, additional `left` and `right` parameters indicating the search range.

## Output Format

- `BinarySearch` and `BinarySearchRecursive` should return the index of the target if found, or -1 if not found.
- `FindInsertPosition` should return the index where the target should be inserted to maintain sorted order.

## Requirements

1. All functions must implement the binary search algorithm, which has O(log n) time complexity.
2. The arrays can be assumed to be sorted in ascending order.
3. `BinarySearchRecursive` must use recursion to solve the problem.
4. If multiple occurrences of the target exist, return the index of any occurrence.

## Sample Input and Output

### Sample Input 1

```
BinarySearch([]int{1, 3, 5, 7, 9}, 5)
```

### Sample Output 1

```
2
```

### Sample Input 2

```
BinarySearch([]int{1, 3, 5, 7, 9}, 6)
```

### Sample Output 2

```
-1
```

### Sample Input 3

```
BinarySearchRecursive([]int{1, 3, 5, 7, 9}, 7, 0, 4)
```

### Sample Output 3

```
3
```

### Sample Input 4

```
FindInsertPosition([]int{1, 3, 5, 7, 9}, 6)
```

### Sample Output 4

```
3
```

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-21/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required functions.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-21/` directory:

```bash
go test -v
``` 