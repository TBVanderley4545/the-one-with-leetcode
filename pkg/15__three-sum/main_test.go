package threesum_test

import (
	"reflect"
	"testing"

	primaryPackage "github.com/TBVanderley4545/the-one-with-leetcode/pkg/15__three-sum"
)

func TestThreeSumSorted(t *testing.T) {
	input1 := []int{-1, 0, 1, 2, -1, -4}
	expected1 := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	actual1 := primaryPackage.ThreeSumSorted(input1)

	input2 := []int{0, 1, 1}
	expected2 := [][]int{}
	actual2 := primaryPackage.ThreeSumSorted(input2)

	input3 := []int{0, 0, 0}
	expected3 := [][]int{{0, 0, 0}}
	actual3 := primaryPackage.ThreeSumSorted(input3)

	input4 := []int{-2, 0, 0, 2, 2}
	expected4 := [][]int{{-2, 0, 2}}
	actual4 := primaryPackage.ThreeSumSorted(input4)

	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Expected the output to be %v, but received %v.", expected1, actual1)
	}

	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Expected the output to be %v, but received %v.", expected2, actual2)
	}

	if !reflect.DeepEqual(expected3, actual3) {
		t.Errorf("Expected the output to be %v, but received %v.", expected3, actual3)
	}

	if !reflect.DeepEqual(expected4, actual4) {
		t.Errorf("Expected the output to be %v, but received %v.", expected4, actual4)
	}
}
