/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (31.38%)
 * Likes:    5517
 * Dislikes: 1425
 * Total Accepted:    936.2K
 * Total Submissions: 3M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Example:
 *
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	ptr1, ptr2 := l1, l2
	var prev1 *ListNode
	carry := 0

	for ptr1 != nil && ptr2 != nil {
		value := ptr1.Val + ptr2.Val + carry
		ptr1.Val = value % 10
		carry = value / 10
		prev1 = ptr1
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	if ptr1 == nil && ptr2 != nil {
		prev1.Next = ptr2
	}

	for ptr := prev1.Next; ptr != nil && carry > 0; ptr = ptr.Next {
		value := ptr.Val + carry
		ptr.Val = value % 10
		carry = value / 10
		prev1 = ptr
	}

	if carry > 0 {
		prev1.Next = &ListNode{
			Val: carry,
		}
	}

	return l1
}
