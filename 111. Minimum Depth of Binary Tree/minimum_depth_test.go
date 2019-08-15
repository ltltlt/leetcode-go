package minimum_depth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func TestMinDepth(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	assert.Equal(t, 2, minDepth(tree))
	assert.Equal(t, 0, minDepth(nil))
	assert.Equal(t, 1, minDepth(&TreeNode{Val: 1}))
	assert.Equal(t, 2, minDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))
}
