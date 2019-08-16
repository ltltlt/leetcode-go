package iterator_test

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func pushAllLeft(root *TreeNode, stack []*TreeNode) []*TreeNode {
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	return stack
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{pushAllLeft(root, nil)}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	length := len(this.stack)
	tail := this.stack[length-1]
	this.stack = pushAllLeft(tail.Right, this.stack[:length-1])
	return tail.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
