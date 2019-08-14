package insert_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	ptr := &root
	for *ptr != nil {
		if (*ptr).Val > val {
			ptr = &(*ptr).Left
		} else {
			ptr = &(*ptr).Right
		}
	}
	*ptr = &TreeNode{Val: val}
	return root
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := inorder(root.Left)
	result = append(result, root.Val)
	return append(result, inorder(root.Right)...)
}
