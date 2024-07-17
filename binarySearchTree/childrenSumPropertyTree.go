package binarySearchTree

func (n *Node) ChildrenSum() bool {
	if n == nil {
		return true
	}

	if n.left == nil && n.right == nil {
		return true
	}

	sum := 0
	if n.left != nil {
		sum += n.left.Value
	}
	if n.right != nil {
		sum += n.right.Value
	}

	return n.Value == sum && n.left.ChildrenSum() && n.right.ChildrenSum()
}
