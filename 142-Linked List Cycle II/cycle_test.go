package cycle_test

import (
	"testing"

	. "github.com/ltltlt/leetcode-util/listutil"
	"github.com/stretchr/testify/assert"
)

func detectCycle(head *ListNode) *ListNode {
	ptr1, ptr2 := head, head
	for {
		if ptr2 == nil || ptr2.Next == nil {
			return nil
		}
		ptr2 = ptr2.Next.Next
		ptr1 = ptr1.Next

		if ptr1 == ptr2 {
			break
		}
	}
	ptr1 = head
	// same speed, ptr1 and ptr2 will meet at the cicle head
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}

func TestHasCycle(t *testing.T) {
	list := NewList([]int{3, 2, 0, -4})
	list.Next.Next.Next.Next = list.Next
	assert.Equal(t, list.Next, detectCycle(list))

	list = NewList([]int{1})
	assert.Nil(t, detectCycle(list))

	list = NewList([]int{1, 2})
	list.Next.Next = list
	assert.Equal(t, list, detectCycle(list))

	list = NewList([]int{2, 2})
	assert.Nil(t, detectCycle(list))
}
