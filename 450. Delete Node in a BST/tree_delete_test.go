package tree_delete_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 两种方法, 因为要修改找到节点在树中的状态, 所以需要至少其中一个信息:　这个节点父节点的指针，或`这个节点父节点中指向这个节点的指针`的指针；我用了后一个
// 这种方法在一些语言中没法用，比如python或java

func deleteNode(root *TreeNode, key int) *TreeNode {
	// passing root's pointer but not root so we can change this local variable
	pointer := findNodePointer(&root, key)
	if pointer == nil {
		return root
	}
	if (*pointer).Left == nil {
		*pointer = (*pointer).Right
	} else if (*pointer).Right == nil {
		*pointer = (*pointer).Left
	} else {
		successorPointer := findSuccessorPointer(*pointer)
		(*pointer).Val = (*successorPointer).Val
		*successorPointer = (*successorPointer).Right
	}

	return root
}

func findSuccessorPointer(node *TreeNode) **TreeNode {
	if node.Right == nil {
		panic("node.Right should not be nil")
	}
	left := &(node.Right)
	for ; (*left).Left != nil; left = &(*left).Left {
	}
	return left
}

func findNodePointer(rootPtr **TreeNode, key int) **TreeNode {
	root := *rootPtr
	if root == nil {
		return nil
	}
	if root.Val == key {
		return rootPtr
	}
	// following loop follows this invariant: root.Val never equals to key
	for {
		if root == nil {
			return nil
		}
		if root.Left != nil && root.Left.Val == key {
			return &root.Left
		} else if root.Right != nil && root.Right.Val == key {
			return &root.Right
		}
		if root.Val > key {
			root = root.Left
		} else if root.Val < key {
			root = root.Right
		}
	}
}
