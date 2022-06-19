package longestsubstring_test

import (
	"testing"

	"github.com/TBVanderley4545/the-one-with-leetcode/pkg/3__longest-substring"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	input1 := "abcabcbb"
	expected1 := 3
	actual1 := longestsubstring.LengthOfLongestSubstring(input1)

	input2 := "bbbbb"
	expected2 := 1
	actual2 := longestsubstring.LengthOfLongestSubstring(input2)

	input3 := "pwwkew"
	expected3 := 3
	actual3 := longestsubstring.LengthOfLongestSubstring(input3)

	input4 := ""
	expected4 := 0
	actual4 := longestsubstring.LengthOfLongestSubstring(input4)

	input5 := " "
	expected5 := 1
	actual5 := longestsubstring.LengthOfLongestSubstring(input5)

	input6 := "au"
	expected6 := 2
	actual6 := longestsubstring.LengthOfLongestSubstring(input6)

	input7 := "abba"
	expected7 := 2
	actual7 := longestsubstring.LengthOfLongestSubstring(input7)

	if expected1 != actual1 {
		t.Errorf("For an input of %s, the expected output is: %d. The actual output was %d",
			input1,
			expected1,
			actual1)
	}

	if expected2 != actual2 {
		t.Errorf("For an input of %s, the expected output is: %d. The actual output was %d",
			input2,
			expected2,
			actual2)
	}

	if expected3 != actual3 {
		t.Errorf("For an input of %s, the expected output is: %d. The actual output was %d",
			input3,
			expected3,
			actual3)
	}

	if expected4 != actual4 {
		t.Errorf("For an input of %s, the expected output is: %d. The actual output was %d",
			input4,
			expected4,
			actual4)
	}

	if expected5 != actual5 {
		t.Errorf("For an input of %s, the expected output is: %d. The actual output was %d",
			input5,
			expected5,
			actual5)
	}

	if expected6 != actual6 {
		t.Errorf("For an input of %s, the expected output is: %d. The actual output was %d",
			input6,
			expected6,
			actual6)
	}

	if expected7 != actual7 {
		t.Errorf("For an input of %s, the expected output is: %d. The actual output was %d",
			input7,
			expected7,
			actual7)
	}
}
