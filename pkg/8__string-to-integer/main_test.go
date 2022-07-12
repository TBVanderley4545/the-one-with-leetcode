package stringtointeger_test

import (
	"testing"

	primary "github.com/TBVanderley4545/the-one-with-leetcode/pkg/8__string-to-integer"
)

func TestMyAtoiImproved(t *testing.T) {
	input1 := "42"
	expected1 := 42
	actual1 := primary.MyAtoiImproved(input1)

	input2 := "   -42"
	expected2 := -42
	actual2 := primary.MyAtoiImproved(input2)

	input3 := "4193 with words"
	expected3 := 4193
	actual3 := primary.MyAtoiImproved(input3)

	input4 := "+1"
	expected4 := 1
	actual4 := primary.MyAtoiImproved(input4)

	input5 := "20000000000000000000"
	expected5 := 2147483647
	actual5 := primary.MyAtoiImproved(input5)

	input6 := "        +1003508d0457"
	expected6 := 1003508
	actual6 := primary.MyAtoiImproved(input6)

	input7 := "-+12"
	expected7 := 0
	actual7 := primary.MyAtoiImproved(input7)

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

	if actual7 != expected7 {
		t.Errorf("Expected the output to be %d, but received %d.", expected7, actual7)
	}
}
