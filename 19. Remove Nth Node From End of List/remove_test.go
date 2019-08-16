package remove_test

type ListNode struct {
	Val int
	Next *ListNode
}

// as leetcode suggested, Maintain two pointers and update one with a delay of n steps.

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return head
	}

	ptr1 := head

	for i:=0; i<n; i++ {
		ptr1 = ptr1.Next
	}

	if ptr1 == nil {
		return head.Next
	}

	ptr2 := head

	for ptr1.Next != nil {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	ptr2.Next = ptr2.Next.Next

	return head
}
