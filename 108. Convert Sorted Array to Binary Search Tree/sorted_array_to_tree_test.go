package build_tree_test

import (
	. "github.com/ltltlt/leetcode-util/treeutil"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	switch len(nums) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: nums[0]}
	default:
		middle := len(nums) / 2
		return &TreeNode{
			Val:   nums[middle],
			Left:  sortedArrayToBST(nums[:middle]),
			Right: sortedArrayToBST(nums[middle+1:]),
		}
	}
}
