package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLenAndTail(head *ListNode) (int, *ListNode) {
	current := head
	var l int
	var prev *ListNode
	for current != nil {
		prev = current
		l++
		current = current.Next
	}
	return l, prev
}

// cost O(n) time and O(1) space
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, tailA := getLenAndTail(headA)
	lenB, tailB := getLenAndTail(headB)

	if tailA != tailB {
		return nil
	}

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func TestIntersectionNode(t *testing.T) {
	l1 := &ListNode{
		Next: &ListNode{
			Next: &ListNode{},
		},
	}
	l2 := &ListNode{
		Next: &ListNode{
			Next: l1,
		},
	}
	l3 := &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: l1,
			},
		},
	}
	assert.Equal(t, l1, getIntersectionNode(l2, l3))
	assert.Equal(t, l1, getIntersectionNode(l1, l2))
}
