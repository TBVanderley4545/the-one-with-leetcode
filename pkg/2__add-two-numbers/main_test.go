package addtwonumbers_test

import (
	"testing"

	addtwonumbers "github.com/TBVanderley4545/the-one-with-leetcode/pkg/2__add-two-numbers"
)

func createListFromArray(numArr []int) *addtwonumbers.ListNode {
	if len(numArr) == 0 {
		return nil
	}

	if len(numArr) == 1 {
		return &addtwonumbers.ListNode{Val: numArr[0]}
	}

	return &addtwonumbers.ListNode{
		Val:  numArr[0],
		Next: createListFromArray(numArr[1:])}
}

func TestAddTwoNumbers(t *testing.T) {
	l1List := createListFromArray([]int{2, 4, 3})
	l2List := createListFromArray([]int{5, 6, 4})
	expectedList := createListFromArray([]int{7, 0, 8})

	actualList := addtwonumbers.AddTwoNumbers(l1List, l2List)

	actualHead := actualList
	actualN1 := actualHead.Next
	actualTail := actualN1.Next

	expectedHead := expectedList
	expectedN1 := expectedHead.Next
	expectedTail := expectedN1.Next

	if actualHead.Val != expectedHead.Val {
		t.Errorf("Expected to see %d, but received %d.",
			expectedHead.Val,
			actualHead.Val)
	}

	if actualN1.Val != expectedN1.Val {
		t.Errorf("Expected to see %d, but received %d.",
			expectedN1.Val,
			actualN1.Val)
	}

	if actualTail.Val != expectedTail.Val {
		t.Errorf("Expected to see %d, but received %d.",
			expectedTail.Val,
			actualTail.Val)
	}

	if actualTail.Next != nil {
		t.Errorf("Expected the tail to be nil, but saw %v",
			actualTail.Next)
	}
}
