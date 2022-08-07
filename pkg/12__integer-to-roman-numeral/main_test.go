package integertoromannumeral_test

import (
	"testing"

	RomanNum "github.com/TBVanderley4545/the-one-with-leetcode/pkg/12__integer-to-roman-numeral"
)

func TestIntToRoman(t *testing.T) {
	input1 := 3
	expected1 := "III"
	actual1 := RomanNum.IntToRoman(input1)

	input2 := 58
	expected2 := "LVIII"
	actual2 := RomanNum.IntToRoman(input2)

	input3 := 1994
	expected3 := "MCMXCIV"
	actual3 := RomanNum.IntToRoman(input3)

	if actual1 != expected1 {
		t.Errorf("Expected the output to be %s, but received %s.", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("Expected the output to be %s, but received %s.", expected2, actual2)
	}

	if actual3 != expected3 {
		t.Errorf("Expected the output to be %s, but received %s.", expected3, actual3)
	}
}
