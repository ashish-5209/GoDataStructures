package binarySearchTree

import "fmt"

func (n *Node) RightView(level int) {
	if n == nil {
		return
	}

	if maxLevel < level {
		fmt.Print(n.Value, " ")
		maxLevel = level
	}

	n.right.RightView(level + 1)
	n.left.RightView(level + 1)
}
