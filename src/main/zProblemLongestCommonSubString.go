package main

import (
	"fmt"
)

// LongestCommonSubstring finds and returns the longest contiguous
// shared sequence of characters between two strings.
func LongestCommonSubstring(s1, s2 string) string {
	// Convert to runes to safely handle multi-byte UTF-8 characters
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	m := len(runes1)
	n := len(runes2)

	// Create a 2D DP table initialized to 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLength := 0
	endIndex := 0 // Tracks the ending position in runes1

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If characters match, extend the previous diagonal sequence
			if runes1[i-1] == runes2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

				// Track the global maximum length and position
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endIndex = i
				}
			} else {
				// Characters don't match; sequence resets to 0
				dp[i][j] = 0
			}
		}
	}

	// If no common substring exists, return an empty string
	if maxLength == 0 {
		return ""
	}

	// Slice and convert the rune array back to a string
	return string(runes1[endIndex-maxLength : endIndex])
}

func main() {
	str1 := "ABABC"
	str2 := "BABCA"

	result := LongestCommonSubstring(str1, str2)
	fmt.Printf("Longest Common Substring: %q\n", result)
	// Output: "BABC"
}
