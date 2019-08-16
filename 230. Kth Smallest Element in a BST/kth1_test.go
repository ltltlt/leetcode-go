package kth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// better solution: find node while iterate

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest1(root *TreeNode, k int) int {
	stack := pushAllLeft(root, nil)

	for len(stack) > 0 {
		l := len(stack)
		tail := stack[l-1]
		stack = stack[:l-1]
		k--
		if k == 0 {
			return tail.Val
		}
		stack = pushAllLeft(tail.Right, stack)
	}
	return -1
}

func pushAllLeft(root *TreeNode, stack []*TreeNode) []*TreeNode {
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	return stack
}

func TestKth(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	assert.Equal(t, 3, kthSmallest1(root, 3))
}