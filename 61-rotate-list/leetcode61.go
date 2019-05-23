package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var length int
	var tail *ListNode
	for cur := head; cur != nil; cur, length = cur.Next, length+1 {
		tail = cur
	}
	pos := (length - (k % length)) % length
	if pos == 0 {
		return head
	}
	var cur, prev *ListNode
	cur = head
	for i := 0; i < pos; i++ {
		cur, prev = cur.Next, cur
	}
	tail.Next = head
	prev.Next = nil
	return cur
}
