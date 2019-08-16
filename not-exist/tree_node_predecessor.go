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

func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, preorder(root.Left)...)
	result = append(result, preorder(root.Right)...)
	return result
}

func preorderIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	var result []int
	for len(stack) > 0 {
		l := len(stack)
		tail := stack[l-1]
		stack = stack[:l-1]
		result = append(result, tail.Val)
		if tail.Right != nil {
			stack = append(stack, tail.Right)
		}
		if tail.Left != nil {
			stack = append(stack, tail.Left)
		}
	}
	return result
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := inorder(root.Left)
	result = append(result, root.Val)
	return append(result, inorder(root.Right)...)
}

func inorderIterative(root *TreeNode) []int {
	stack := pushAllLeft(root, nil)
	result := make([]int, 0, len(stack))

	for len(stack) > 0 {
		l := len(stack)
		tail := stack[l-1]
		stack = stack[:l-1]
		result = append(result, tail.Val)
		stack = pushAllLeft(tail.Right, stack)
	}
	return nil
}

func pushAllLeft(root *TreeNode, stack []*TreeNode) []*TreeNode {
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	return stack
}

func postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := postorder(root.Left)
	result = append(result, postorder(root.Right)...)
	result = append(result, root.Val)
	return result
}

func postorderIterative(root *TreeNode) []int {
	// TODO
}
