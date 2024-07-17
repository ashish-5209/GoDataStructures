package binarySearchTree

import "fmt"

var maxLevel int = 0

func (n *Node) LeftView(level int) {
	if n == nil {
		return
	}

	if maxLevel < level {
		fmt.Print(n.Value, " ")
		maxLevel = level
	}

	n.left.LeftView(level + 1)
	n.right.LeftView(level + 1)
}
