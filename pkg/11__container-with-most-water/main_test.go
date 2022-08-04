package containerwithmostwater_test

import (
	"testing"

	watermodule "github.com/TBVanderley4545/the-one-with-leetcode/pkg/11__container-with-most-water"
)

func TestMaxArea(t *testing.T) {
	input1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected1 := 49

	input2 := []int{1, 1}
	expected2 := 1

	actual1 := watermodule.MaxArea(input1)
	actual2 := watermodule.MaxArea(input2)

	if actual1 != expected1 {
		t.Errorf("Expected the output to be %d, but received %d.", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("Expected the output to be %d, but received %d.", expected2, actual2)
	}
}
