package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(list1, list2 *ListNode) *ListNode {
	dummyHead := ListNode{}

	l1Node := list1
	l2Node := list2
	current := &dummyHead
	carry := 0

	for l1Node != nil || l2Node != nil {
		sum := carry

		if l1Node != nil {
			sum += l1Node.Val
			l1Node = l1Node.Next
		}

		if l2Node != nil {
			sum += l2Node.Val
			l2Node = l2Node.Next
		}

		carry = sum / 10
		nextNode := ListNode{Val: sum % 10}

		current.Next = &nextNode

		current = current.Next
	}

	if carry > 0 {
		nextNode := ListNode{Val: carry}
		current.Next = &nextNode
	}

	return dummyHead.Next
}
