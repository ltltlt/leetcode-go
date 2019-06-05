package reverse_list_test

import (
	"testing"

	. "github.com/ltltlt/leetcode-util/listutil"
	"github.com/stretchr/testify/assert"
)

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var prev *ListNode // stores previous element of reversed list
	current := head
	for i := 1; i < m && current != nil; i++ {
		prev = current
		current = current.Next
	}

	headr := current // stores header of reversed list(when not reversed)

	var prev1 *ListNode // store previous of current processing element
	for i := m; i <= n && current != nil; i++ {
		tmp := current.Next
		current.Next = prev1
		prev1 = current
		current = tmp
	}
	if prev != nil {
		prev.Next = prev1
	}
	if headr != nil {
		headr.Next = current
	}

	if m <= 1 {
		return prev1
	}
	return head
}

func TestReverseBetween(t *testing.T) {
	l := NewList([]int{1, 2, 3, 4, 5})

	l = reverseBetween(l, 2, 4)
	ListMatchSlice(t, l, []int{1, 4, 3, 2, 5})

	assert.Nil(t, reverseBetween(nil, 2, 4))

	l = NewList([]int{5})

	assert.Equal(t, l, reverseBetween(l, 2, 4))
	assert.Equal(t, l, reverseBetween(l, 0, 0))

	l = NewList([]int{3, 5})
	l = reverseBetween(l, 1, 2)
	ListMatchSlice(t, l, []int{5, 3})
}
