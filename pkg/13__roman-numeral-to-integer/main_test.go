package romannumeraltointeger_test

import (
	"testing"

	RomanNum "github.com/TBVanderley4545/the-one-with-leetcode/pkg/13__roman-numeral-to-integer"
)

func TestRomanToInt(t *testing.T) {
	input1 := "III"
	expected1 := 3
	actual1 := RomanNum.RomanToInt(input1)

	input2 := "LVIII"
	expected2 := 58
	actual2 := RomanNum.RomanToInt(input2)

	input3 := "MCMXCIV"
	expected3 := 1994
	actual3 := RomanNum.RomanToInt(input3)

	if actual1 != expected1 {
		t.Errorf("Expected the output to be %d, but received %d.", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("Expected the output to be %d, but received %d.", expected2, actual2)
	}

	if actual3 != expected3 {
		t.Errorf("Expected the output to be %d, but received %d.", expected3, actual3)
	}
}
