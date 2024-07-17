package binarySearchTree

import (
	"container/list"
	"fmt"
)

func (n *Node) DiagonalView() {
	if n == nil {
		return
	}

	l := list.New()
	l.PushBack(n)

	for l.Len() > 0 {
		size := l.Len()
		for size > 0 {
			temp := l.Front().Value.(*Node)
			l.Remove(l.Front())
			for temp != nil {
				fmt.Print(temp.Value, " ")
				if temp.left != nil {
					l.PushBack(temp.left)
				}
				temp = temp.right
			}
			size -= 1
		}
		fmt.Println()
	}
}
