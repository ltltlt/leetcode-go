package palindrome_test

import (
	"testing"

	. "github.com/ltltlt/leetcode-util/listutil"
	"github.com/stretchr/testify/assert"
)

// microsoft china interview
// solution is simple, cost O(n) times and O(1) space
// find middle, reverse middle to end, compare start to middle with middle to end
// finally, reverse middle to end back

func isPalindrome(head *ListNode) bool {
	size := getSize(head)
	if size < 2 {
		return true
	}

	half := size / 2
	middle := head
	for i := 0; i < half; i++ {
		middle = middle.Next
	}
	if half*2 != size {
		middle = middle.Next
	}
	middle = reverse(middle)
	defer reverse(middle)

	ptr1 := head
	ptr2 := middle
	for i := 0; i < half; i++ {
		if ptr1.Val != ptr2.Val {
			return false
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return true
}

func getSize(head *ListNode) int {
	var size int
	for current := head; current != nil; current = current.Next {
		size++
	}
	return size
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for current := head; current != nil; {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	return prev
}

func TestIsPalindrome(t *testing.T) {
	for i, c := range []struct {
		input  []int
		output bool
	}{
		{[]int{1, 2}, false},
		{[]int{1, 2, 2, 1}, true},
		{nil, true},
		{[]int{1}, true},
	} {
		assert.Equalf(t, c.output, isPalindrome(NewList(c.input)), "case %d fail", i)
	}
}
