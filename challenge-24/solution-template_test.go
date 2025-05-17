package main

import (
	"testing"
)

func TestDPLongestIncreasingSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Basic example", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"Multiple possible LIS", []int{0, 1, 0, 3, 2, 3}, 4},
		{"All same numbers", []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{"Non-trivial example", []int{4, 10, 4, 3, 8, 9}, 3},
		{"Empty array", []int{}, 0},
		{"Single element", []int{5}, 1},
		{"Decreasing order", []int{5, 4, 3, 2, 1}, 1},
		{"Increasing order", []int{1, 2, 3, 4, 5}, 5},
		{"Complex example", []int{3, 10, 2, 1, 20}, 3},
		{"Another complex example", []int{50, 3, 10, 7, 40, 80}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DPLongestIncreasingSubsequence(tt.nums)
			if result != tt.expected {
				t.Errorf("DPLongestIncreasingSubsequence(%v) = %d, expected %d",
					tt.nums, result, tt.expected)
			}
		})
	}
}

func TestOptimizedLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Basic example", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"Multiple possible LIS", []int{0, 1, 0, 3, 2, 3}, 4},
		{"All same numbers", []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{"Non-trivial example", []int{4, 10, 4, 3, 8, 9}, 3},
		{"Empty array", []int{}, 0},
		{"Single element", []int{5}, 1},
		{"Decreasing order", []int{5, 4, 3, 2, 1}, 1},
		{"Increasing order", []int{1, 2, 3, 4, 5}, 5},
		{"Complex example", []int{3, 10, 2, 1, 20}, 3},
		{"Another complex example", []int{50, 3, 10, 7, 40, 80}, 4},
		{"Large example", generateLargeTestCase(1000), 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OptimizedLIS(tt.nums)
			if result != tt.expected {
				t.Errorf("OptimizedLIS(%v) = %d, expected %d",
					truncateForDisplay(tt.nums), result, tt.expected)
			}
		})
	}
}

func TestGetLISElements(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		expectedLength int
		isIncreasing   bool // Check if the returned sequence is strictly increasing
	}{
		{"Basic example", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4, true},
		{"Multiple possible LIS", []int{0, 1, 0, 3, 2, 3}, 4, true},
		{"All same numbers", []int{7, 7, 7, 7, 7, 7, 7}, 1, true},
		{"Non-trivial example", []int{4, 10, 4, 3, 8, 9}, 3, true},
		{"Empty array", []int{}, 0, true},
		{"Single element", []int{5}, 1, true},
		{"Decreasing order", []int{5, 4, 3, 2, 1}, 1, true},
		{"Increasing order", []int{1, 2, 3, 4, 5}, 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetLISElements(tt.nums)

			// Check the length of the result
			if len(result) != tt.expectedLength {
				t.Errorf("GetLISElements(%v) returned %v with length %d, expected length %d",
					tt.nums, result, len(result), tt.expectedLength)
			}

			// Check if the result is a valid LIS
			if len(result) > 0 && tt.isIncreasing {
				isValid := true
				for i := 1; i < len(result); i++ {
					if result[i] <= result[i-1] {
						isValid = false
						break
					}
				}
				if !isValid {
					t.Errorf("GetLISElements(%v) returned %v, which is not strictly increasing",
						tt.nums, result)
				}
			}

			// Check if the result is a valid subsequence
			if len(result) > 0 && !isValidSubsequence(tt.nums, result) {
				t.Errorf("GetLISElements(%v) returned %v, which is not a valid subsequence",
					tt.nums, result)
			}
		})
	}
}

func TestExampleCases(t *testing.T) {
	// Example 1
	nums1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expected1 := 4

	if result := DPLongestIncreasingSubsequence(nums1); result != expected1 {
		t.Errorf("Example 1: DPLongestIncreasingSubsequence(%v) = %d, expected %d",
			nums1, result, expected1)
	}

	// Example 2
	nums2 := []int{0, 1, 0, 3, 2, 3}
	expected2 := 4

	if result := OptimizedLIS(nums2); result != expected2 {
		t.Errorf("Example 2: OptimizedLIS(%v) = %d, expected %d",
			nums2, result, expected2)
	}

	// Example 3
	nums3 := []int{10, 9, 2, 5, 3, 7, 101, 18}

	result3 := GetLISElements(nums3)
	if len(result3) != 4 {
		t.Errorf("Example 3: GetLISElements(%v) = %v with length %d, expected length 4",
			nums3, result3, len(result3))
	}

	// Example 4
	nums4 := []int{7, 7, 7, 7, 7, 7, 7}
	expected4 := 1

	if result := DPLongestIncreasingSubsequence(nums4); result != expected4 {
		t.Errorf("Example 4: DPLongestIncreasingSubsequence(%v) = %d, expected %d",
			nums4, result, expected4)
	}
}

// Helper function to check if a sequence is a valid subsequence of another
func isValidSubsequence(nums []int, sub []int) bool {
	if len(sub) == 0 {
		return true
	}
	if len(nums) == 0 {
		return false
	}

	subIdx := 0
	for _, num := range nums {
		if num == sub[subIdx] {
			subIdx++
			if subIdx == len(sub) {
				return true
			}
		}
	}
	return false
}

// Helper function to generate a large test case
func generateLargeTestCase(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			result[i] = i
		} else {
			result[i] = size - i
		}
	}
	return result
}

// Helper function to truncate arrays for better error displays
func truncateForDisplay(nums []int) []int {
	if len(nums) <= 10 {
		return nums
	}
	return append(nums[:5], nums[len(nums)-5:]...)
}

// Benchmark the performance of the different implementations
func BenchmarkLIS(b *testing.B) {
	benchCases := []struct {
		name string
		nums []int
	}{
		{"Small case", []int{10, 9, 2, 5, 3, 7, 101, 18}},
		{"Medium case", []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}},
		{"Large case", generateLargeTestCase(1000)},
	}

	for _, bc := range benchCases {
		b.Run("DP-"+bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DPLongestIncreasingSubsequence(bc.nums)
			}
		})

		b.Run("Optimized-"+bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				OptimizedLIS(bc.nums)
			}
		})

		b.Run("GetElements-"+bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetLISElements(bc.nums)
			}
		})
	}
}
