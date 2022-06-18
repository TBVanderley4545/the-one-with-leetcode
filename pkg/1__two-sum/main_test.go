package twosum_test

import (
	"reflect"
	"testing"

	"github.com/TBVanderley4545/the-one-with-leetcode/pkg/1__two-sum"
)

func TestTwoSum(t *testing.T) {
	input1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}

	input2 := []int{3, 2, 4}
	target2 := 6
	expected2 := []int{1, 2}

	input3 := []int{3, 3}
	target3 := 6
	expected3 := []int{0, 1}

	actual1 := twosum.TwoSum(input1, target1)
	actual2 := twosum.TwoSum(input2, target2)
	actual3 := twosum.TwoSum(input3, target3)

	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Error in first data set. Expected to see %v, but received %v.", expected1, actual1)
	}

	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Error in second data set. Expected to see %v, but received %v.", expected2, actual2)
	}

	if !reflect.DeepEqual(expected3, actual3) {
		t.Errorf("Error in third data set. Expected to see %v, but received %v.", expected3, actual3)
	}
}
