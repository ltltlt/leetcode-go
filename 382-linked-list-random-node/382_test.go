package solution_test

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	count int64
	head  *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	var count int64
	current := head
	for current != nil {
		current = current.Next
		count = count + 1
	}
	return Solution{count, head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	random := rand.Int63n(this.count)
	current := this.head
	for i := int64(0); i < random; i++ {
		current = current.Next
	}
	return current.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
