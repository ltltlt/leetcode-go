package reorder_list_test

import (
	"testing"

	. "github.com/ltltlt/leetcode-util/listutil"
)

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	ptr1, ptr2 := head, head
	// find middle
	for ptr2 != nil {
		ptr2 = ptr2.Next
		if ptr2 == nil {
			break
		}
		ptr2 = ptr2.Next
		if ptr2 == nil {
			break
		}
		ptr1 = ptr1.Next
	}

	// reverse middle to end
	var prev *ListNode
	current := ptr1.Next
	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	// 1->2->3->6->5->4 => 1->6->2->5->3->4
	l1 := head
	l2 := prev
	for l2 != nil {
		tmp1 := l1.Next
		l1.Next = l2
		tmp2 := l2.Next
		l2.Next = tmp1
		l1 = tmp1
		l2 = tmp2
	}
	l1.Next = nil
}

func TestReorderList(t *testing.T) {
	l := NewList([]int{1, 2, 3, 4, 5, 6})

	reorderList(l)

	ListMatchSlice(t, l, []int{1, 6, 2, 5, 3, 4})

	l = NewList([]int{1})
	reorderList(l)
	ListMatchSlice(t, l, []int{1})
}
