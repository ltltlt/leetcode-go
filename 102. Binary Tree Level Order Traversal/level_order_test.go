package level_order_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var levels [][]int

	q := []*TreeNode{root}
	for len(q) != 0 {
		var newq []*TreeNode
		level := make([]int, len(q))
		for i, element := range q {
			level[i] = element.Val
			if element.Left != nil {
				newq = append(newq, element.Left)
			}
			if element.Right != nil {
				newq = append(newq, element.Right)
			}
		}
		q = newq
		levels = append(levels, level)
	}
	return levels
}

func TestLevelOrder(t *testing.T) {
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
				Val: 7,
			},
		},
	}

	r := levelOrder(tree)
	assert.ElementsMatch(t, []int{3}, r[0])
	assert.ElementsMatch(t, []int{9, 20}, r[1])
	assert.ElementsMatch(t, []int{15, 7}, r[2])
}
