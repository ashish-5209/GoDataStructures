package binarySearchTree

import "math"

func (n *Node) MaxOfTree() int {
	if n == nil {
		return math.MinInt
	}
	return max(n.Value, max(n.left.MaxOfTree(), n.right.MaxOfTree()))
}
