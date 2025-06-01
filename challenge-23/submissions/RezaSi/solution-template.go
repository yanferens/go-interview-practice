package main

import (
	"fmt"
)

func main() {
	// Sample texts and patterns
	testCases := []struct {
		text    string
		pattern string
	}{
		{"ABABDABACDABABCABAB", "ABABCABAB"},
		{"AABAACAADAABAABA", "AABA"},
		{"GEEKSFORGEEKS", "GEEK"},
		{"AAAAAA", "AA"},
	}

	// Test each pattern matching algorithm
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Text: %s\n", tc.text)
		fmt.Printf("Pattern: %s\n", tc.pattern)

		// Test naive pattern matching
		naiveResults := NaivePatternMatch(tc.text, tc.pattern)
		fmt.Printf("Naive Pattern Match: %v\n", naiveResults)

		// Test KMP algorithm
		kmpResults := KMPSearch(tc.text, tc.pattern)
		fmt.Printf("KMP Search: %v\n", kmpResults)

		// Test Rabin-Karp algorithm
		rkResults := RabinKarpSearch(tc.text, tc.pattern)
		fmt.Printf("Rabin-Karp Search: %v\n", rkResults)

		fmt.Println("------------------------------")
	}
}

// NaivePatternMatch performs a brute force search for pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func NaivePatternMatch(text, pattern string) []int {
    startingIndices := make([]int, 0)
    if len(pattern) == 0 {
        return startingIndices
    }

	for i := 0; i < len(text); i++ {
	    if i + len(pattern) > len(text) {
	        break
	    }

	    if pattern == text[i: i + len(pattern)] {
	        startingIndices = append(startingIndices, i)
	    }
	}
	
	return startingIndices
}

// KMPSearch implements the Knuth-Morris-Pratt algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func KMPSearch(text, pattern string) []int {
    matches := []int{}
    
    // Handle edge cases
    if len(pattern) == 0 || len(text) < len(pattern) {
        return matches
    }
    
    n := len(text)
    m := len(pattern)
    
    // Preprocess the pattern
    lps := computeLPSArray(pattern)
    
    i := 0 // Index for text
    j := 0 // Index for pattern
    
    for i < n {
        // Current characters match, move both pointers forward
        if pattern[j] == text[i] {
            i++
            j++
        }
        
        // Found a complete match
        if j == m {
            matches = append(matches, i-j)
            // Use lps to shift pattern for next match
            j = lps[j-1]
        } else if i < n && pattern[j] != text[i] {
            // Mismatch after j matches
            if j != 0 {
                // Use lps to shift pattern
                j = lps[j-1]
            } else {
                // No match found, move to next character in text
                i++
            }
        }
    }
    
    return matches
}

func computeLPSArray(pattern string) []int {
    m := len(pattern)
    lps := make([]int, m)
    
    // Length of the previous longest prefix suffix
    length := 0
    i := 1
    
    // The loop calculates lps[i] for i = 1 to m-1
    for i < m {
        if pattern[i] == pattern[length] {
            length++
            lps[i] = length
            i++
        } else {
            // This is the tricky part
            if length != 0 {
                length = lps[length-1]
                // Note: We do not increment i here
            } else {
                lps[i] = 0
                i++
            }
        }
    }
    
    return lps
}

// RabinKarpSearch implements the Rabin-Karp algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func RabinKarpSearch(text, pattern string) []int {
    matches := []int{}
    
    // Handle edge cases
    if len(pattern) == 0 || len(text) < len(pattern) {
        return matches
    }
    
    n := len(text)
    m := len(pattern)
    
    // Large prime number to avoid hash collisions
    prime := 101
    
    // Base value for the hash function
    base := 256
    
    // Hash value for pattern and initial window
    patternHash := 0
    windowHash := 0
    
    // Highest power of base that we need
    h := 1
    for i := 0; i < m-1; i++ {
        h = (h * base) % prime
    }
    
    // Calculate initial hash values
    for i := 0; i < m; i++ {
        patternHash = (base*patternHash + int(pattern[i])) % prime
        windowHash = (base*windowHash + int(text[i])) % prime
    }
    
    // Slide the pattern over text one by one
    for i := 0; i <= n-m; i++ {
        // Check if hash values match
        if patternHash == windowHash {
            // Verify the match character by character
            match := true
            for j := 0; j < m; j++ {
                if text[i+j] != pattern[j] {
                    match = false
                    break
                }
            }
            if match {
                matches = append(matches, i)
            }
        }
        
        // Calculate hash value for next window
        if i < n-m {
            windowHash = (base*(windowHash-int(text[i])*h) + int(text[i+m])) % prime
            
            // Ensure we only have positive hash values
            if windowHash < 0 {
                windowHash += prime
            }
        }
    }
    
    return matches
}
