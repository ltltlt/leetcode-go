package middle_test

import (
	"testing"

	. "github.com/ltltlt/leetcode-util/listutil"
	"github.com/stretchr/testify/assert"
)

func middleNode(head *ListNode) *ListNode {
	ptr1, ptr2 := head, head
	for ptr2 != nil {
		ptr2 = ptr2.Next
		if ptr2 == nil {
			break
		}
		ptr2 = ptr2.Next
		ptr1 = ptr1.Next
	}
	return ptr1
}

func TestMiddle(t *testing.T) {
	l := NewList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 3, middleNode(l).Val)

	l = NewList([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 4, middleNode(l).Val)

	l = NewList([]int{1})
	assert.Equal(t, 1, middleNode(l).Val)

	l = nil
	assert.Nil(t, middleNode(l))
}
