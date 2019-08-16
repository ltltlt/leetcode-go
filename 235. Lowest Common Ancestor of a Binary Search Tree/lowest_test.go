package lowest_test

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}

	path1 := getPath(root, p)
	path2 := getPath(root, q)

	if len(path1) < len(path2) {
		path1, path2 = path2, path1
	}

	for i, node := range path1 {
		if len(path2) <= i || path2[i] != node {
			return path1[i-1]
		}
	}
	return nil
}

func getPath(root *TreeNode, p *TreeNode) []*TreeNode {
	var result []*TreeNode
	current := root

	for {
		if current == nil {
			return nil
		}
		result = append(result, current)
		if current == p {
			break
		}
		if current.Val > p.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return result
}