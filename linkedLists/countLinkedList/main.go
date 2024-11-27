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

func (l *LinkedList) Insert(val int) {
	newNode := &Node{data: val}
	currHead := l.head
	l.head = newNode
	l.head.next = currHead
	l.count++
}

func (l LinkedList) PrintListData() {
	for l.count > 0 && l.head != nil {
		fmt.Print(l.head.data, " ")
		l.head = l.head.next
		l.count--
	}
	fmt.Println()
}

func main() {
	head := LinkedList{}
	head.Insert(10)
	head.Insert(20)
	head.Insert(30)
	head.Insert(40)
	head.Insert(50)
	head.Insert(60)

	head.PrintListData()

	fmt.Println(head.head.data, head.count)
}
