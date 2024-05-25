package main

import (
	"fmt"
	linkedList "gods/LinkedList"
	"gods/heap"
	"gods/stack"
)

func main() {
	dsMap := make(map[int]string)
	dsMap[0] = "Linked List"
	dsMap[1] = "Heap"
	dsMap[2] = "Stack"

	for key, val := range dsMap {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case 0:
		fmt.Println(dsMap[val])
		linkedList.App()
	case 1:
		fmt.Println(dsMap[val])
		heap.App()
	case 2:
		fmt.Println(dsMap[val])
		stack.App()
	}
}
