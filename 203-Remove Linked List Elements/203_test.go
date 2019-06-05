package remove_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for ptr := &head; ptr != nil && *ptr != nil; {
		if (*ptr).Val == val {
			*ptr = (*ptr).Next
		} else {
			ptr = &(*ptr).Next
		}
	}
	return head
}

func newList(elems []int) *ListNode {
	if len(elems) == 0 {
		return nil
	}

	head := &ListNode{Val: elems[0]}
	var prev = head
	for _, elem := range elems[1:] {
		prev.Next = &ListNode{Val: elem}
		prev = prev.Next
	}

	return head
}

func match(t *testing.T, head *ListNode, elems []int) {
	for _, elem := range elems {
		assert.Equal(t, elem, head.Val)
		head = head.Next
	}
	assert.Nil(t, head)
}

func TestRemoveElements(t *testing.T) {
	l := newList([]int{1, 2, 6, 3, 4, 5, 6})
	l = removeElements(l, 6)
	match(t, l, []int{1, 2, 3, 4, 5})

	l = newList([]int{1, 1})
	l = removeElements(l, 1)
	match(t, l, nil)
}
