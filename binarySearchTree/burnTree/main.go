package main

import (
	"fmt"
)

type Node struct {
	Value       int
	left, right *Node
}

func newNode(val int) *Node {
	return &Node{
		Value: val,
	}
}

func (n *Node) printNode() {
	if n == nil {
		return
	}
	n.left.printNode()
	fmt.Print(n.Value, " ")
	n.right.printNode()
}

func (n *Node) BurnTree(timer *int, target int) int {
	if n == nil {
		return 0
	}

	if n.Value == target {
		return -1
	}

	left := n.left.BurnTree(timer, target)
	right := n.right.BurnTree(timer, target)

	if left < 0 {
		*timer = max(*timer, -1*left+right)
		return left - 1
	}

	if right < 0 {
		*timer = max(*timer, -1*right+left)
		return right - 1
	}

	return 1 + max(left, right)

}

func (n *Node) find(node *Node, target int) {
	if n == nil {
		return
	}
	if n.Value == target {
		node = n
		return
	}

	n.left.find(node, target)
	n.right.find(node, target)
}

func (n *Node) height() int {
	if n == nil {
		return 0
	}

	return 1 + max(n.left.height(), n.right.height())
}

func main() {
	head := newNode(1)
	head.left = newNode(2)
	head.left.left = newNode(4)
	head.left.right = newNode(5)
	head.left.right.left = newNode(7)
	head.left.right.right = newNode(8)
	head.right = newNode(3)
	head.right.right = newNode(6)
	head.right.right.right = newNode(9)
	head.right.right.right.right = newNode(10)

	head.printNode()
	fmt.Println()

	timer := 0
	head.BurnTree(&timer, 8)
	var node *Node
	head.find(node, 8)
	h := node.height() - 1
	fmt.Println(max(h, timer))
}
