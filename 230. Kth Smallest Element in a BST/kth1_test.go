package kth1_test

// TODO: complete

import (
	. "github.com/ltltlt/leetcode-util/treeutil"
)

// better solution: find node while iterate

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var stack []*TreeNode

	for 
}


func dfs(root *TreeNode) []*TreeNode {
	stack := []*TreeNode{root}
	var result []*TreeNode
	for len(stack) > 0 {
		last := stack[len(stack)-1]

	}
}