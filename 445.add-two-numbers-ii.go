/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 *
 * https://leetcode.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (50.51%)
 * Likes:    783
 * Dislikes: 104
 * Total Accepted:    96.8K
 * Total Submissions: 191.5K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The most significant digit comes first and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the
 * lists is not allowed.
 *
 *
 *
 * Example:
 *
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * it would be easy if I can reverse list, just reverse l1 and l2,
 * iterate through it, add each, may need carry bit, create a new node...
 * even though we cannot reverse the list, we can still use O(n) space to
 * store backward pointer, so we don't need O(n^2) time to do the carry bit
 * thing
 * this program is over-complexity, because I don't modify origin list at all
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	len1 := getLen(l1)
	len2 := getLen(l2)
	var stack []*ListNode
	longer, shorter := l1, l2
	if len1 < len2 {
		longer, shorter = l2, l1
		len1, len2 = len2, len1
	}
	for i := 0; i < len1-len2; i++ {
		n := clone(longer)
		if len(stack) > 0 {
			stack[len(stack)-1].Next = n
		}
		stack = append(stack, n)
		longer = longer.Next
	}
	var head *ListNode
	for i := 0; i < len2; i++ {
		value := longer.Val + shorter.Val
		longer = longer.Next
		shorter = shorter.Next
		node := &ListNode{
			Val: value % 10,
		}
		l := len(stack) - 1
		if l >= 0 {
			stack[l].Next = node
		}
		stack = append(stack, node)
		for ; l >= 0 && value >= 10; l-- { // value will always be smaller than 20
			value = stack[l].Val + 1
			stack[l].Val = value % 10
		}
		if value >= 10 { // this can cost a lot, luckily, this can happen only once, so I use some tricky to avoid array move
			head = &ListNode{
				Val:  1,
				Next: stack[0],
			}
		}
	}
	if head != nil {
		return head
	}
	return stack[0]
}

func getLen(list *ListNode) int {
	size := 0
	for list != nil {
		size++
		list = list.Next
	}
	return size
}

func clone(list *ListNode) *ListNode {
	l := *list
	return &l
}

