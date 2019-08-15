package path_sum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if isLeaf(root) {
		return sum == root.Val
	}

	newSum := sum - root.Val
	return hasPathSum(root.Left, newSum) || hasPathSum(root.Right, newSum)
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func TestPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
		},
	}

	assert.True(t, hasPathSum(&TreeNode{Val: 22}, 22))
	assert.True(t, hasPathSum(tree, 22))
	assert.False(t, hasPathSum(&TreeNode{Val: 22}, 23))
	assert.False(t, hasPathSum(nil, 22))
	assert.False(t, hasPathSum(nil, 0))
	assert.False(t, hasPathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 1))
}
