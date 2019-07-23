package construct_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (41.94%)
 * Likes:    1849
 * Dislikes: 51
 * Total Accepted:    237.6K
 * Total Submissions: 564.5K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	value := preorder[0]
	pos := mustIndex(inorder, value)
	return &TreeNode{
		Val:   value,
		Left:  buildTree(preorder[1:pos+1], inorder[0:pos]),
		Right: buildTree(preorder[pos+1:], inorder[pos+1:]),
	}
}

func mustIndex(slice []int, elem int) int {
	for pos, e := range slice {
		if e == elem {
			return pos
		}
	}
	panic("not found")
}

func treeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && treeEqual(t1.Left, t2.Left) && treeEqual(t1.Right, t2.Right)
}

func TestBuildTree(t *testing.T) {
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
	tree1 := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	assert.True(t, treeEqual(tree, tree1))
}
