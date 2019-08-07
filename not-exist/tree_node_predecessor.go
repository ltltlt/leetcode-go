package tree_node_predecessor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func findPredecessor(root *TreeNode, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return nil
	}
	if root.Left != nil {
		var right *TreeNode
		for right = root.Left; right.Right != nil; right = right.Right {
		}
		return right
	}
	parent := node.Parent
	for parent != nil && parent.Left == node {
		node = parent
		parent = node.Parent
	}
	return parent
}

func TestFindPredessor(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -1,
		},
	}
	root.Left.Parent = root
	assert.Equal(t, findPredecessor(root, root).Val, -1)
	assert.Nil(t, findPredecessor(root, root.Left))
}
