package binarySearchTree

import (
	"container/list"
	"fmt"
)

func (n *Node) LevelOrderTraversal() {
	if n == nil {
		return
	}

	l := list.New()
	l.PushBack(n)
	l.PushBack(nil) // Marker for the end of the current level

	for l.Len() > 0 {
		curr := l.Front().Value
		l.Remove(l.Front())

		if curr == nil {
			fmt.Println() // End of the current level

			if l.Len() > 0 {
				l.PushBack(nil) // Marker for the next level
			}
			continue
		}
		node := curr.(*Node) // Safe to cast now
		fmt.Print(node.Value, " ")

		if node.left != nil {
			l.PushBack(node.left)
		}
		if node.right != nil {
			l.PushBack(node.right)
		}

	}
}
