package reverse_linked_list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}
	return prev
}

func TestReverseList(t *testing.T) {
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
	l = reverseList(l)
	assert.Equal(t, 5, l.Val)
	assert.Equal(t, 4, l.Next.Val)
	assert.Equal(t, 3, l.Next.Next.Val)
	assert.Equal(t, 2, l.Next.Next.Next.Val)
	assert.Equal(t, 1, l.Next.Next.Next.Next.Val)
	assert.Equal(t, (*ListNode)(nil), l.Next.Next.Next.Next.Next)
}
