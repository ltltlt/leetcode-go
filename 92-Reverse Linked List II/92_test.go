package reverse_list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	l := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3, &ListNode{
					4, &ListNode{
						5, nil,
					},
				},
			},
		},
	}

	l = reverseBetween(l, 2, 4)
	assert.Equal(t, 1, l.Val)
	assert.Equal(t, 4, l.Next.Val)
	assert.Equal(t, 3, l.Next.Next.Val)
	assert.Equal(t, 2, l.Next.Next.Next.Val)
	assert.Equal(t, 5, l.Next.Next.Next.Next.Val)
	assert.Equal(t, (*ListNode)(nil), l.Next.Next.Next.Next.Next)

	assert.Equal(t, (*ListNode)(nil), reverseBetween(nil, 2, 4))

	l = &ListNode{
		Val: 5,
	}

	assert.Equal(t, l, reverseBetween(l, 2, 4))
	assert.Equal(t, l, reverseBetween(l, 0, 0))

	l = &ListNode{
		3, &ListNode{5, nil},
	}
	l = reverseBetween(l, 1, 2)
	assert.Equal(t, 5, l.Val)
	assert.Equal(t, 3, l.Next.Val)
}
