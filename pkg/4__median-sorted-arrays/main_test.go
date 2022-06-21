package mediansortedarrays_test

import (
	"testing"

	"github.com/TBVanderley4545/the-one-with-leetcode/pkg/4__median-sorted-arrays"
)

func outputTestResults(t *testing.T, expected, actual float64) {
	if expected != actual {
		t.Errorf("Expected to see %f, but received %f", expected, actual)
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3, 6}
	nums2 := []int{2, 4, 8}
	expected1 := 3.5

	nums3 := []int{10, 30, 60}
	nums4 := []int{20, 40, 70, 90}
	expected2 := 40.0

	nums5 := []int{}
	nums6 := []int{1}
	expected3 := 1.0

	nums7 := []int{6}
	nums8 := []int{8}
	expected4 := 7.0

	nums9 := []int{1, 2}
	nums10 := []int{3, 4}
	expected5 := 2.5

	actual1 := mediansortedarrays.FindMedianSortedArrays(nums1, nums2)
	actual2 := mediansortedarrays.FindMedianSortedArrays(nums3, nums4)
	actual3 := mediansortedarrays.FindMedianSortedArrays(nums5, nums6)
	actual4 := mediansortedarrays.FindMedianSortedArrays(nums7, nums8)
	actual5 := mediansortedarrays.FindMedianSortedArrays(nums9, nums10)

	outputTestResults(t, expected1, actual1)
	outputTestResults(t, expected2, actual2)
	outputTestResults(t, expected3, actual3)
	outputTestResults(t, expected4, actual4)
	outputTestResults(t, expected5, actual5)
}
