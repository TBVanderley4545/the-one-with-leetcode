package reverseinteger_test

import (
	"testing"

	"github.com/TBVanderley4545/the-one-with-leetcode/pkg/7__reverse-integer"
)

func TestReverse(t *testing.T) {

	input1 := 123
	expected1 := 321
	actual1 := reverseinteger.Reverse(input1)

	input2 := -123
	expected2 := -321
	actual2 := reverseinteger.Reverse(input2)

	input3 := 120
	expected3 := 21
	actual3 := reverseinteger.Reverse(input3)

	input4 := -2147483412
	expected4 := -2143847412
	actual4 := reverseinteger.Reverse(input4)

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
}
