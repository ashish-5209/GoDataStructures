package main

import (
	"container/list"
	"fmt"
)

func main() {
	n := 50
	l := list.New()
	l.PushBack("1")
	i := 0
	for i < n {
		curr := l.Front()
		l.Remove(curr)
		fmt.Print(curr.Value, " ")
		l.PushBack(curr.Value.(string) + "0")
		l.PushBack(curr.Value.(string) + "1")
		i += 1
	}
	fmt.Println()
}
