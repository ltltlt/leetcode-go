package validate_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	valid, _, _ := recursive(root)
	return valid
}

func recursive(root *TreeNode) (bool, int, int) {
	var valid bool
	var lmin, lmax, rmin, rmax int
	lmin, rmax = root.Val, root.Val
	if root.Left != nil {
		valid, lmin, lmax = recursive(root.Left)
		if !valid || lmax >= root.Val {
			return false, 0, 0
		}
	}
	if root.Right != nil {
		valid, rmin, rmax = recursive(root.Right)
		if !valid || rmin <= root.Val {
			return false, 0, 0
		}
	}
	return true, lmin, rmax
}
