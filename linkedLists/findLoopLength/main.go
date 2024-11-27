package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head  *Node
	count int
}

type LinkedListI interface {
	Insert(int)
	FindLoopLength() int
	CreateLoop(int, int)
}

func (l *LinkedList) Insert(val int) {
	newNode := &Node{data: val}
	currHead := l.head
	l.head = newNode
	l.head.next = currHead
	l.count++
}

func (l *LinkedList) CreateLoop(val1, val2 int) {
	node1, node2 := l.head, l.head

	for node1 != nil {
		if node1.data == val1 {
			break
		}
		node1 = node1.next

	}
	for node2 != nil {
		if node2.data == val2 {
			break
		}
		node2 = node2.next

	}
	node2.next = node1

}

func (l *LinkedList) FindLoopLength() int {
	slow, fast := l.head, l.head

	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			slow = slow.next
			res := 1
			for slow != fast {
				res++
				slow = slow.next
			}
			return res
		}
	}
	return 0
}

func main() {
	var ll LinkedListI
	ll = &LinkedList{}
	ll.Insert(45)
	ll.Insert(58)
	ll.Insert(90)
	ll.Insert(39)
	ll.Insert(21)
	ll.Insert(10)
	ll.Insert(33)
	ll.Insert(19)
	ll.Insert(14)
	ll.Insert(25)
	ll.CreateLoop(33, 45)

	fmt.Println(ll.FindLoopLength())
}
