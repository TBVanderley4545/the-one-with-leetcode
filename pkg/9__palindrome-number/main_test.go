package palindromenumber_test

import (
	"testing"

	primary "github.com/TBVanderley4545/the-one-with-leetcode/pkg/9__palindrome-number"
)

func TestIsPalindrome(t *testing.T) {
	input1 := 121
	expected1 := true
	actual1 := primary.IsPalindrome(input1)

	input2 := -121
	expected2 := false
	actual2 := primary.IsPalindrome(input2)

	input3 := 10
	expected3 := false
	actual3 := primary.IsPalindrome(input3)

	if actual1 != expected1 {
		t.Errorf("Expected the output to be %t, but received %t.", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("Expected the output to be %t, but received %t.", expected2, actual2)
	}

	if actual3 != expected3 {
		t.Errorf("Expected the output to be %t, but received %t.", expected3, actual3)
	}
}
