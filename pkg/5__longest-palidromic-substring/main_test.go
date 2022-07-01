package longestpalidromicsubstring_test

import (
	"testing"

	"github.com/TBVanderley4545/the-one-with-leetcode/pkg/5__longest-palidromic-substring"
)

func TestLongestPalindrome(t *testing.T) {
	input1 := "babad"
	expected1 := "bab"
	expected1Alt := "aba"
	actual1 := longestpalidromicsubstring.LongestPalindrome(input1)

	input2 := "cbbd"
	expected2 := "bb"
	actual2 := longestpalidromicsubstring.LongestPalindrome(input2)

	input3 := "1"
	expected3 := "1"
	actual3 := longestpalidromicsubstring.LongestPalindrome(input3)

	input4 := "abacdfgdcaba"
	expected4 := "aba"
	actual4 := longestpalidromicsubstring.LongestPalindrome(input4)

	input5 := "caba"
	expected5 := "aba"
	actual5 := longestpalidromicsubstring.LongestPalindrome(input5)

	if actual1 != expected1 && actual1 != expected1Alt {
		t.Errorf("Expected the longest palindrome to be %s, but received %s.", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("Expected the longest palindrome to be %s, but received %s.", expected2, actual2)
	}

	if actual3 != expected3 {
		t.Errorf("Expected the longest palindrome to be %s, but received %s.", expected3, actual3)
	}

	if actual4 != expected4 {
		t.Errorf("Expected the longest palindrome to be %s, but received %s.", expected4, actual4)
	}

	if actual5 != expected5 {
		t.Errorf("Expected the longest palindrome to be %s, but received %s.", expected5, actual5)
	}
}
