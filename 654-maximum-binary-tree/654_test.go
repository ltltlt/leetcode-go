package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMaxIndex(nums []int, i, j int) int {
	if i >= j {
		panic("i >= j")
	}
	var max = -1
	var index = -1
	for ; i < j; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index
}

func recursive(nums []int, i, j int) *TreeNode {
	idx := findMaxIndex(nums, i, j)
	node := TreeNode{
		Val: nums[idx],
	}
	if i < idx {
		node.Left = recursive(nums, i, idx)
	}
	if idx+1 < j {
		node.Right = recursive(nums, idx+1, j)
	}
	return &node
}

// apparent solution: find max value, use it divide nums into left and right
// part, and call it on left and right recursively
// the problem is how to find max value

// easy way is find max each time, this lead time complexity to O(n logn);

// another way is use a segment tree, this costs to build a segment tree is O(n logn)
// and cost extra space O(n), following steps cost only O(log n * log n);

// we can also use two dimension array to store the max value for each range,
// but this cost O(n^2) to preprocess and O(n^2) space
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	return recursive(nums, 0, len(nums))
}

// Another solution, one iteration,
// this cost only k*n(k is a constant) time, since each one loop is needed to push and pop one times
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var stack []*TreeNode
	for _, num := range nums {
		node := &TreeNode{Val: num}
		j := len(stack) - 1
		for ; j >= 0 && stack[j].Val < num; j-- {
		}
		if j+1 < len(stack) {
			node.Left = stack[j+1]
		}
		if j >= 0 {
			stack[j].Right = node
		}
		stack = stack[:j+1]
		stack = append(stack, node)
	}
	if len(stack) > 0 {
		return stack[0]
	}
	return nil
}

func TestTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	tree := constructMaximumBinaryTree(nums)
	assert.Equal(t, 6, tree.Val)
	assert.Equal(t, 3, tree.Left.Val)
	assert.Equal(t, 2, tree.Left.Right.Val)
	assert.Equal(t, 1, tree.Left.Right.Right.Val)
	assert.Equal(t, 5, tree.Right.Val)
	assert.Equal(t, 0, tree.Right.Left.Val)
}
