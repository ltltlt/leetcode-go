package cycle_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n) time complexity(at most loop k+m times, m is the cycle length, k is the non-cycle length),
// and O(1) space complexity
func hasCycle(head *ListNode) bool {
	ptr1, ptr2 := head, head
	for ptr2 != nil {
		ptr1 = ptr1.Next
		if ptr2.Next == nil {
			break
		}
		ptr2 = ptr2.Next.Next
		if ptr1 == ptr2 { // there is no way both ptr1 and ptr2 be nil, because ptr2 is ahead of ptr1
			return true
		}
	}
	return false
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
	assert.True(t, hasCycle(list))

	list = newList([]int{1})
	assert.False(t, hasCycle(list))

	list = newList([]int{1, 2})
	list.Next.Next = list
	assert.True(t, hasCycle(list))

	list = newList([]int{2, 2})
	assert.False(t, hasCycle(list))
}
