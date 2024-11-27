package main

import (
	"fmt"
	"math"
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

func (n *Node) PathSum(val *int) int {
	if n == nil {
		return 0
	}

	if n.left == nil && n.right == nil {
		return n.Value
	}

	left := n.left.PathSum(val)
	right := n.right.PathSum(val)

	if n.left != nil && n.right != nil {
		*val = max(*val, n.Value+left+right)
		return n.Value + max(left, right)
	}

	if n.left != nil {
		return n.Value + left
	}

	return n.Value + right

}

func main() {
	root := newNode(10)
	root.left = newNode(2)
	root.left.left = newNode(20)
	root.left.right = newNode(1)
	root.right = newNode(10)
	root.right.right = newNode(-25)
	root.right.right.left = newNode(3)
	root.right.right.right = newNode(4)
	val := math.MinInt
	res := root.PathSum(&val)

	if root.left != nil && root.right != nil {
		fmt.Println(res)
	} else {
		fmt.Println(max(val, res))
	}
}
