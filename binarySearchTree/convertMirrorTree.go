package binarySearchTree

func (n *Node) MirrorTree() {
	if n == nil {
		return
	}
	n.left.MirrorTree()
	n.right.MirrorTree()
	n.left, n.right = n.right, n.left
}
