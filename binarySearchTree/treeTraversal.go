package binarySearchTree

import "fmt"

func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.left.InOrderTraversal()
	fmt.Print(n.Value, " ")
	n.right.InOrderTraversal()
}

func (n *Node) PreOrderTraversal() {
	if n == nil {
		return
	}
	fmt.Print(n.Value, " ")
	n.left.PreOrderTraversal()
	n.right.PreOrderTraversal()
}

func (n *Node) PostOrderTraversal() {
	if n == nil {
		return
	}
	n.left.PostOrderTraversal()
	n.right.PostOrderTraversal()
	fmt.Print(n.Value, " ")
}
