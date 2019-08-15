package balanced_test

import (
	"testing"

	. "github.com/ltltlt/leetcode-util/treeutil"
	"github.com/stretchr/testify/assert"
)

// balanced tree: a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
// this problem is weird, the depth seems like the max depth

func isBalanced(root *TreeNode) bool {
	isBalanced1, _ := recursive(root)
	return isBalanced1
}

func recursive(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	isBalanced1, depth1 := recursive(root.Left)
	isBalanced2, depth2 := recursive(root.Right)
	return isBalanced1 && isBalanced2 && diff(depth1, depth2) <= 1, max(depth1, depth2) + 1
}

func diff(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestIsBalanced(t *testing.T) {
	tree := BuildTree([]int{3, 9, 20, NULL, NULL, 15, 7})
	assert.True(t, isBalanced(tree))
	assert.True(t, isBalanced(nil))
	tree = BuildTree([]int{1, 2, 2, 3, 3, NULL, NULL, 4, 4})
	assert.False(t, isBalanced(tree))

	tree = BuildTree([]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, NULL, NULL, 5, 5})
	assert.True(t, isBalanced(tree))
}
