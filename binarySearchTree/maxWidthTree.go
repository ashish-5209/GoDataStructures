package binarySearchTree

import "container/list"

func (n *Node) MaxWidth() int {
	if n == nil {
		return 0
	}

	maxLen := 0

	l := list.New()
	l.PushBack(n)

	for l.Len() > 0 {
		currLen := l.Len()
		maxLen = max(maxLen, currLen)
		for currLen > 0 {
			currLen -= 1
			curr := l.Front().Value.(*Node)
			l.Remove(l.Front())

			if curr.left != nil {
				l.PushBack(curr.left)
			}
			if curr.right != nil {
				l.PushBack(curr.right)
			}
		}
	}

	return maxLen
}
