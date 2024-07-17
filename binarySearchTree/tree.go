package binarySearchTree

import "fmt"

type Node struct {
	Value int
	left  *Node
	right *Node
}

func newNode(val int) *Node {
	return &Node{Value: val}
}

func (n *Node) Insert(value int) {
	if n == nil {
		return
	}

	if value < n.Value {
		if n.left == nil {
			n.left = newNode(value)
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = newNode(value)
		} else {
			n.right.Insert(value)
		}
	}
}

func (n *Node) Search(value int) *Node {
	if n == nil || n.Value == value {
		return n
	}
	if value < n.Value {
		return n.left.Search(value)
	}

	return n.right.Search(value)
}

func InsertSearch() {
	root := newNode(10)
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Print("In-Order Traversal: ")
	root.InOrderTraversal()
	fmt.Println()

	fmt.Print("Pre-Order Traversal: ")
	root.PreOrderTraversal()
	fmt.Println()

	fmt.Print("Post-Order Traversal: ")
	root.PostOrderTraversal()
	fmt.Println()

	searchValue := 7
	foundNode := root.Search(searchValue)

	if foundNode != nil {
		fmt.Printf("Found node with value: %d\n", foundNode.Value)
	} else {
		fmt.Printf("Node with value %d not found\n", searchValue)
	}

	fmt.Println("Height of BST:", root.Height())

	root.PrintKthNodes(0)
	fmt.Println()

	root.LevelOrderTraversal()

	fmt.Println(root.SizeOfTree())

	fmt.Println(root.MaxOfTree())
	root.LeftView(1)
	fmt.Println()
	maxLevel = 0
	root.RightView(1)
	fmt.Println()

	root.DiagonalView()

	fmt.Println(root.ChildrenSum())

	fmt.Println(root.CheckBalancedTree())

	fmt.Println(root.MaxWidth())
}
