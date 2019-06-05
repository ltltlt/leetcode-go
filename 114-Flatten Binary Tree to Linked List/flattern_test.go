package flattern_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(node *TreeNode, stack []*TreeNode) (*TreeNode, []*TreeNode) {
	if node == nil {
		if l := len(stack); l > 0 {
			return stack[l-1], stack
		}
		return nil, nil
	}

	if node.Right != nil {
		stack = append(stack, node.Right)
	}
	node.Right, stack = recursive(node.Left, stack)
	node.Left = nil
	return node, stack
}

func flatten(root *TreeNode) {
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		_, stack = recursive(root, stack)
	}
}

func TestFlattern(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)
	assert.Equal(t, 1, root.Val)
	assert.Equal(t, 2, root.Right.Val)
	assert.Equal(t, 3, root.Right.Right.Val)
	assert.Equal(t, 4, root.Right.Right.Right.Val)
	assert.Equal(t, 5, root.Right.Right.Right.Right.Val)
	assert.Equal(t, 6, root.Right.Right.Right.Right.Right.Val)
	assert.Nil(t, root.Right.Right.Right.Right.Right.Right)
}
