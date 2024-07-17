package binarySearchTree

import "math"

func (n *Node) CheckBalancedTree() (bool, float64) {
	if n == nil {
		return true, 0
	}

	leftBool, lh := n.left.CheckBalancedTree()
	rightBool, rh := n.right.CheckBalancedTree()

	diff := math.Abs(lh-rh) <= 1

	height := max(lh, rh) + 1

	if leftBool && rightBool && diff {
		return true, height
	}
	return false, height
}
