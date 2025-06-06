package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	// TODO: Implement the function
	byteSlice := []byte(s)
	for i :=0 ; i < len(s)/2 ; i ++ {
	    byteSlice[i], byteSlice[len(s)-i-1] = byteSlice[len(s)-i-1], byteSlice[i] 
	}
	return string(byteSlice)
}
