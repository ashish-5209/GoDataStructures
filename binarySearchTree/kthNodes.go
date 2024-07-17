package binarySearchTree

import "fmt"

func (n *Node) PrintKthNodes(k int) {
	if n == nil {
		return
	}

	if k == 0 {
		fmt.Print(n.Value, " ")
		return
	}
	n.left.PrintKthNodes(k - 1)
	n.right.PrintKthNodes(k - 1)
}
