package cycle_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func newList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func TestHasCycle(t *testing.T) {
	list := newList([]int{3, 2, 0, -4})
	list.Next.Next.Next.Next = list.Next
	assert.Equal(t, list.Next, detectCycle(list))

	list = newList([]int{1})
	assert.Nil(t, detectCycle(list))

	list = newList([]int{1, 2})
	list.Next.Next = list
	assert.Equal(t, list, detectCycle(list))

	list = newList([]int{2, 2})
	assert.Nil(t, detectCycle(list))
}
