package kth_test

// plain solution: find minimun node, then call k-1 times to find successor

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	node := findMin(root)
	for ; k > 1; k-- {
		node = findSucessor(root, node)
	}

	return node.Val
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root
	for ; left.Left != nil; left = left.Left {
	}
	return left
}

func findSucessor(root, node *TreeNode) *TreeNode {
	if node.Right != nil {
		return findMin(node.Right)
	}

	path := []*TreeNode{root}
	for {
		last := path[len(path)-1]
		if last.Val == node.Val {
			break
		} else if last.Val > node.Val {
			path = append(path, last.Left)
		} else {
			path = append(path, last.Right)
		}
	}

	var previous *TreeNode
	var i int
	for i = len(path) - 1; i > 0; i-- {
		previous = path[i-1]
		current := path[i]
		if previous.Left == current {
			break
		}
	}
	if i == 0 {
		return nil
	}
	return previous
}
