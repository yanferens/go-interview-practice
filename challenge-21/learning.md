# Learning Materials for Binary Search

## Understanding Binary Search

Binary search is a fundamental algorithm in computer science that efficiently finds an element in a sorted array by repeatedly dividing the search interval in half. It's a perfect example of the divide-and-conquer algorithmic paradigm.

### How Binary Search Works

1. Start with the middle element of the array
2. If the target value equals the middle element, we're done
3. If the target value is less than the middle element, search the left half
4. If the target value is greater than the middle element, search the right half
5. Repeat until the element is found or the search space is empty

Binary search has a time complexity of O(log n), which is much more efficient than linear search (O(n)) for large datasets.

### Iterative Implementation

Here's how to implement binary search iteratively in Go:

```go
func binarySearch(arr []int, target int) int {
    left := 0
    right := len(arr) - 1
    
    for left <= right {
        mid := left + (right - left) / 2  // Avoid potential overflow
        
        if arr[mid] == target {
            return mid  // Found the target
        }
        
        if arr[mid] < target {
            left = mid + 1  // Target is in the right half
        } else {
            right = mid - 1  // Target is in the left half
        }
    }
    
    return -1  // Target not found
}
```

### Recursive Implementation

Binary search can also be implemented recursively:

```go
func binarySearchRecursive(arr []int, target int, left int, right int) int {
    // Base case: element not found
    if left > right {
        return -1
    }
    
    // Calculate mid point
    mid := left + (right - left) / 2
    
    // Check if target is at mid
    if arr[mid] == target {
        return mid
    }
    
    // If target is smaller, search in left half
    if arr[mid] > target {
        return binarySearchRecursive(arr, target, left, mid - 1)
    }
    
    // If target is greater, search in right half
    return binarySearchRecursive(arr, target, mid + 1, right)
}
```

### Finding Insertion Position

A common variant of the binary search problem is to find the position where an element should be inserted to maintain the sorted order:

```go
func findInsertPosition(arr []int, target int) int {
    left := 0
    right := len(arr) - 1
    
    // Handle empty array or target larger than all elements
    if len(arr) == 0 || target > arr[right] {
        return len(arr)
    }
    
    // Handle target smaller than all elements
    if target < arr[0] {
        return 0
    }
    
    for left <= right {
        mid := left + (right - left) / 2
        
        if arr[mid] == target {
            return mid  // Target value already exists
        }
        
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    // At this point, right < left and target should be inserted at index left
    return left
}
```

## Binary Search Edge Cases

When implementing binary search, be mindful of these common edge cases:

1. **Empty array**: Check if the array is empty before searching
2. **Single element array**: Make sure your algorithm works for arrays with just one element
3. **Target smaller than all elements**: Handle the case when the target is smaller than the smallest element
4. **Target larger than all elements**: Handle the case when the target is larger than the largest element
5. **Duplicate elements**: Decide how to handle multiple occurrences of the target value
6. **Integer overflow**: When calculating the middle index, use `mid := left + (right - left) / 2` instead of `mid := (left + right) / 2` to avoid integer overflow for large arrays

## Binary Search Applications

Binary search is used in many real-world applications:

1. **Database Systems**: For indexing and searching records
2. **Debugging**: For finding bugs in large codebases (binary bug search)
3. **Machine Learning**: For hyperparameter tuning
4. **Library Functions**: Used in standard library functions like `sort.Search` in Go
5. **Computer Graphics**: For ray tracing and collision detection algorithms

## Binary Search Variants

There are several variants of binary search:

1. **First occurrence**: Find the first occurrence of an element if duplicates exist
2. **Last occurrence**: Find the last occurrence of an element if duplicates exist
3. **Closest element**: Find the element closest to the target value
4. **Rotated sorted array**: Find an element in a sorted array that has been rotated
5. **2D binary search**: Binary search in a 2D sorted matrix

## Further Reading

1. [Binary Search Algorithm (GeeksforGeeks)](https://www.geeksforgeeks.org/binary-search/)
2. [Binary Search in Go (GoLang Docs)](https://pkg.go.dev/sort#Search)
3. [Divide and Conquer Algorithms (Khan Academy)](https://www.khanacademy.org/computing/computer-science/algorithms/merge-sort/a/divide-and-conquer-algorithms)
4. [Advanced Binary Search Techniques (Topcoder)](https://www.topcoder.com/community/competitive-programming/tutorials/binary-search/)
5. [Binary Search Cookbook (Codeforces)](https://codeforces.com/blog/entry/9901) 