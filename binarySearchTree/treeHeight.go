package binarySearchTree

func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	return max(n.left.Height(), n.right.Height()) + 1
}
