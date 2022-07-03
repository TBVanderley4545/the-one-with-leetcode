package zigzagconversion_test

import (
	"testing"

	conversion "github.com/TBVanderley4545/the-one-with-leetcode/pkg/6__zigzag-conversion"
)

func TestConvert(t *testing.T) {
	input1 := "PAYPALISHIRING"
	expected1 := "PAHNAPLSIIGYIR"
	actual1 := conversion.Convert(input1, 3)

	input2 := "PAYPALISHIRING"
	expected2 := "PINALSIGYAHRPI"
	actual2 := conversion.Convert(input2, 4)

	input3 := "A"
	expected3 := "A"
	actual3 := conversion.Convert(input3, 1)

	input4 := "AB"
	expected4 := "AB"
	actual4 := conversion.Convert(input4, 1)

	if actual1 != expected1 {
		t.Errorf("Expected the output to be %s, but received %s.", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("Expected the output to be %s, but received %s.", expected2, actual2)
	}

	if actual3 != expected3 {
		t.Errorf("Expected the output to be %s, but received %s.", expected3, actual3)
	}

	if actual4 != expected4 {
		t.Errorf("Expected the output to be %s, but received %s.", expected4, actual4)
	}
}
