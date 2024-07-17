package binarySearchTree

func (n *Node) SizeOfTree() int {
	if n == nil {
		return 0
	}

	return 1 + n.left.SizeOfTree() + n.right.SizeOfTree()
}
