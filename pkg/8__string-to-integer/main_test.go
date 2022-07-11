package stringtointeger_test

import (
	"testing"

	primary "github.com/TBVanderley4545/the-one-with-leetcode/pkg/8__string-to-integer"
)

func TestMyAtoi(t *testing.T) {
	input1 := "42"
	expected1 := 42
	actual1 := primary.MyAtoi(input1)

	input2 := "   -42"
	expected2 := -42
	actual2 := primary.MyAtoi(input2)

	input3 := "4193 with words"
	expected3 := 4193
	actual3 := primary.MyAtoi(input3)

	input4 := "+1"
	expected4 := 1
	actual4 := primary.MyAtoi(input4)

	input5 := "20000000000000000000"
	expected5 := 2147483647
	actual5 := primary.MyAtoi(input5)

	input6 := "        +1003508d0457"
	expected6 := 1003508
	actual6 := primary.MyAtoi(input6)

	if actual1 != expected1 {
		t.Errorf("Expected the output to be %d, but received %d.", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("Expected the output to be %d, but received %d.", expected2, actual2)
	}

	if actual3 != expected3 {
		t.Errorf("Expected the output to be %d, but received %d.", expected3, actual3)
	}

	if actual4 != expected4 {
		t.Errorf("Expected the output to be %d, but received %d.", expected4, actual4)
	}

	if actual5 != expected5 {
		t.Errorf("Expected the output to be %d, but received %d.", expected5, actual5)
	}

	if actual6 != expected6 {
		t.Errorf("Expected the output to be %d, but received %d.", expected6, actual6)
	}
}
