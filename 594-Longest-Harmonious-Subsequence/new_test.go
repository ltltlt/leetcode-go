package leetcode594

type node struct {
	count       int
	left, right *node
}

func (n *node) leftIncrement() {
	if n.left == nil {
		n.left = new(node)
	}
	n.left.count += 1
}

func (n *node) rightIncrement() {
	if n.right == nil {
		n.right = new(node)
	}
	n.right.count += 1
}

func findLHS1(nums []int) int {
	var m = make(map[int]*node, len(nums))
	for _, num := range nums {
		if n, ok := m[num]; ok {
			n.count += 1
		} else {
			m[num] = &node{count: 1}
		}
		if n, ok := m[num-1]; ok {
			n.leftIncrement()
		}
		if n, ok := m[num+1]; ok {
			n.rightIncrement()
		}
	}
	var longest = 0
	for _, n := range m {
		if n.left != nil {
			longest = max(longest, n.left.count+n.count)
		}
		if n.right != nil {
			longest = max(longest, n.right.count+n.count)
		}
	}
	return longest
}
