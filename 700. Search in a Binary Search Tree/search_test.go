package search_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 24ms
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

// 28ms
func searchBSTIterative(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root
}
