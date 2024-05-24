package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Insert(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d		", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *LinkedList) Delete(val int) {
	if l.length == 0 {
		return
	}

	previousToDelete := l.head
	if l.head.data == val {
		l.head = l.head.next
		l.length--
	}
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func InsertDelete() {
	myList := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	node4 := &Node{data: 11}
	node5 := &Node{data: 7}
	node6 := &Node{data: 2}
	myList.Insert(node1)
	myList.Insert(node2)
	myList.Insert(node3)
	myList.Insert(node4)
	myList.Insert(node5)
	myList.Insert(node6)
	myList.PrintListData()
	myList.Delete(100)
	myList.PrintListData()
	myList.Delete(2)
	myList.PrintListData()
	emptyList := LinkedList{}
	emptyList.Delete(10)
}
