package longestcommonprefix_test

import (
	"testing"

	primaryPackage "github.com/TBVanderley4545/the-one-with-leetcode/pkg/14__longest-common-prefix"
)

func TestLongestCommonPrefix(t *testing.T) {
	input1 := []string{"flower", "flow", "flight"}
	expected1 := "fl"
	actual1 := primaryPackage.LongestCommonPrefix(input1)

	input2 := []string{"dog", "racecar", "car"}
	expected2 := ""
	actual2 := primaryPackage.LongestCommonPrefix(input2)

	if expected1 != actual1 {
		t.Errorf("Expected the output to be %s, but received %s.", expected1, actual1)
	}

	if expected2 != actual2 {
		t.Errorf("Expected the output to be %s, but received %s.", expected2, actual2)
	}
}
