package main

import (
	"fmt"
)

func main() {
	// Example slice for testing
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Test FindMax
	max := FindMax(numbers)
	fmt.Printf("Maximum value: %d\n", max)

	// Test RemoveDuplicates
	unique := RemoveDuplicates(numbers)
	fmt.Printf("After removing duplicates: %v\n", unique)

	// Test ReverseSlice
	reversed := ReverseSlice(numbers)
	fmt.Printf("Reversed: %v\n", reversed)

	// Test FilterEven
	evenOnly := FilterEven(numbers)
	fmt.Printf("Even numbers only: %v\n", evenOnly)
}

// FindMax returns the maximum value in a slice of integers.
// If the slice is empty, it returns 0.
func FindMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }

    maxNumber := numbers[0]
	for _, number := range numbers {
	    maxNumber = max(maxNumber, number)
	}
	
	return maxNumber
}

// RemoveDuplicates returns a new slice with duplicate values removed,
// preserving the original order of elements.
func RemoveDuplicates(numbers []int) []int {
    mapper := make(map[int]bool, 0)
    uniqNumbers := make([]int, 0)
    
    for _, number := range(numbers) {
        if _, exist := mapper[number]; !exist {
            uniqNumbers = append(uniqNumbers, number)
            mapper[number] = true
        }
    }
    
	return uniqNumbers
}

// ReverseSlice returns a new slice with elements in reverse order.
func ReverseSlice(slice []int) []int {
    reversedSlice := make([]int, len(slice), len(slice))

	for i := 0; i < (len(slice) + 1) / 2; i++ {
	    reversedSlice[i], reversedSlice[len(slice) - i - 1] = slice[len(slice) - i - 1], slice[i]
	}
	
	return reversedSlice
}

// FilterEven returns a new slice containing only the even numbers
// from the original slice.
func FilterEven(numbers []int) []int {
	filteredEven := make([]int, 0)
	
	for _, number := range(numbers){
	    if number % 2 == 0 {
	        filteredEven = append(filteredEven, number)
	    }
	}

	return filteredEven
}
