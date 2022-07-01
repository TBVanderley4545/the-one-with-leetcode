package longestpalidromicsubstring

import (
	"math"
	"strings"
)

func LongestPalindrome(s string) string {
	start := 0
	end := 0

	for idx := range s {
		oddPalLen := expandAroundCenter(s, idx, idx)
		evenPalLen := expandAroundCenter(s, idx, idx+1)

		largerPalLen := int(math.Max(float64(oddPalLen), float64(evenPalLen)))

		if largerPalLen > end-start {
			start = idx - (largerPalLen-1)/2

			end = idx + largerPalLen/2
		}
	}

	substring := strings.Join(strings.Split(s, "")[start:end+1], "")

	return substring
}

func expandAroundCenter(s string, left int, right int) int {
	leftPointer, rightPointer := left, right

	for leftPointer >= 0 && rightPointer < len(s) && s[leftPointer] == s[rightPointer] {
		leftPointer--
		rightPointer++
	}

	return rightPointer - leftPointer - 1
}
