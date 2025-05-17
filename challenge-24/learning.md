# Learning Materials for Longest Increasing Subsequence

## Introduction to Dynamic Programming

Dynamic Programming (DP) is a technique for solving problems by breaking them down into simpler subproblems and storing the results of these subproblems to avoid redundant calculations. It's particularly useful for optimization problems where we want to find the maximum or minimum of something.

DP solutions typically have the following characteristics:
1. **Overlapping Subproblems**: The problem can be broken down into subproblems which are reused multiple times.
2. **Optimal Substructure**: The optimal solution to the problem contains optimal solutions to its subproblems.

## The Longest Increasing Subsequence Problem

The Longest Increasing Subsequence (LIS) problem is a classic example of a problem that can be efficiently solved using dynamic programming. Given a sequence of integers, we need to find the length of the longest subsequence such that all elements of the subsequence are sorted in increasing order.

A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

### Examples

- For the sequence `[10, 9, 2, 5, 3, 7, 101, 18]`, the LIS is `[2, 3, 7, 18]` or `[2, 3, 7, 101]` with a length of 4.
- For the sequence `[0, 1, 0, 3, 2, 3]`, the LIS is `[0, 1, 2, 3]` with a length of 4.
- For the sequence `[7, 7, 7, 7, 7, 7, 7]`, the LIS is `[7]` with a length of 1.

## Approach 1: Standard Dynamic Programming (O(n²))

The standard DP approach uses a 1D array `dp` where `dp[i]` represents the length of the longest increasing subsequence ending at index `i`.

### Algorithm

1. Initialize a `dp` array of the same length as the input array, all with value 1 (the minimum LIS length is 1).
2. For each element at position `i`, look at all previous elements at positions `j < i`:
   - If the current element is greater than the previous element (`nums[i] > nums[j]`), we can extend the LIS ending at `j` by including the current element.
   - Update `dp[i] = max(dp[i], dp[j] + 1)`.
3. The maximum value in the `dp` array is the length of the LIS.

### Implementation in Go

```go
func dpLongestIncreasingSubsequence(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // Initialize dp array with 1 (minimum LIS length)
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
    }
    
    maxLength := 1
    
    // Calculate dp[i] for each position
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
        
        // Update the maximum LIS length
        if dp[i] > maxLength {
            maxLength = dp[i]
        }
    }
    
    return maxLength
}
```

### Time and Space Complexity

- **Time Complexity**: O(n²) because we have nested loops that each run for n iterations.
- **Space Complexity**: O(n) for the dp array.

## Approach 2: Optimized Solution Using Binary Search (O(n log n))

We can improve the time complexity to O(n log n) by using a different approach that utilizes binary search.

### Algorithm

1. Maintain an array `tails` where `tails[i]` represents the smallest value at which an increasing subsequence of length `i+1` ends.
2. For each element in the input array:
   - If it's larger than all values in the `tails` array, append it to the end of the array.
   - Otherwise, find the first element in the `tails` array that is greater than or equal to the current element, and replace it.
3. The length of the `tails` array is the length of the LIS.

### Implementation in Go

```go
func optimizedLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // Initialize tails array
    tails := make([]int, 0, len(nums))
    
    for _, num := range nums {
        // If num is larger than all elements in tails, append it
        if len(tails) == 0 || num > tails[len(tails)-1] {
            tails = append(tails, num)
            continue
        }
        
        // Find the first element in tails that is greater than or equal to num
        index := binarySearch(tails, num)
        tails[index] = num
    }
    
    return len(tails)
}

// Binary search to find the first element >= target
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}
```

### Time and Space Complexity

- **Time Complexity**: O(n log n) - We process each element once, and binary search takes O(log n) time.
- **Space Complexity**: O(n) for the tails array.

## Approach 3: Reconstructing the LIS Elements

To find the actual elements of the LIS, not just its length, we need to modify our approach to keep track of the predecessors.

### Using the O(n²) Approach

1. Maintain a `dp` array as before, where `dp[i]` is the length of the LIS ending at index `i`.
2. Also maintain a `prev` array where `prev[i]` is the index of the previous element in the LIS ending at index `i`.
3. After filling the `dp` and `prev` arrays, trace back from the index with the maximum LIS length to construct the actual subsequence.

### Implementation in Go

```go
func getLISElements(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    
    n := len(nums)
    dp := make([]int, n)
    prev := make([]int, n)
    
    // Initialize dp and prev arrays
    for i := range dp {
        dp[i] = 1
        prev[i] = -1 // -1 indicates no predecessor
    }
    
    // Find the LIS length and track predecessors
    maxLength, maxIndex := 1, 0
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] && dp[j]+1 > dp[i] {
                dp[i] = dp[j] + 1
                prev[i] = j
            }
        }
        
        // Update the maximum LIS length and its ending index
        if dp[i] > maxLength {
            maxLength = dp[i]
            maxIndex = i
        }
    }
    
    // Reconstruct the LIS
    lis := make([]int, maxLength)
    for i := maxLength - 1; i >= 0; i-- {
        lis[i] = nums[maxIndex]
        maxIndex = prev[maxIndex]
    }
    
    return lis
}
```

### Time and Space Complexity

- **Time Complexity**: O(n²) for computing the `dp` and `prev` arrays.
- **Space Complexity**: O(n) for the `dp` and `prev` arrays.

## Common Pitfalls and Edge Cases

1. **Empty Array**: Return 0 or an empty slice.
2. **Single Element**: Return 1 or a slice with that element.
3. **All Elements the Same**: The LIS has length 1.
4. **Decreasing Sequence**: The LIS has length 1.
5. **Duplicate Elements**: Be careful about how you handle the `>` comparison. If you want strictly increasing subsequences, use `>`. If you want non-decreasing subsequences, use `>=`.

## Variations and Extensions

1. **Longest Decreasing Subsequence**: Just reverse the comparison or reverse the array first.
2. **Longest Bitonic Subsequence**: A subsequence that first increases, then decreases.
3. **Maximum Sum Increasing Subsequence**: Find the increasing subsequence with the maximum sum.
4. **Longest Common Subsequence**: Find the longest subsequence common to two sequences.

## Applications in Real Life

1. **Stock Market Analysis**: Finding periods of consistent growth.
2. **DNA Sequence Analysis**: Finding common patterns in DNA.
3. **Version Control Systems**: Finding the most efficient way to transform one file into another.
4. **Text Comparison and Diff Tools**: Finding the largest common blocks of text.

## Further Reading

1. [GeeksforGeeks - Longest Increasing Subsequence](https://www.geeksforgeeks.org/longest-increasing-subsequence-dp-3/)
2. [Algorithms: Design and Analysis, Part 1 (Stanford) - Dynamic Programming](https://www.coursera.org/learn/algorithms-divide-conquer)
3. [Introduction to Algorithms (CLRS) - Chapter 15: Dynamic Programming](https://mitpress.mit.edu/books/introduction-algorithms-third-edition)
4. [Competitive Programmer's Handbook - Chapter 7: Dynamic Programming](https://cses.fi/book/book.pdf) 