package main

import (
	"fmt"
	linkedList "gods/LinkedList"
	"gods/binarySearchTree"
	"gods/graph"
	"gods/heap"
	"gods/queue"
	"gods/stack"
)

func main() {
	dsMap := make(map[int]string)
	dsMap[0] = "Linked List"
	dsMap[1] = "Heap"
	dsMap[2] = "Stack"
	dsMap[3] = "Queue"
	dsMap[4] = "Binary Search Tree"
	dsMap[5] = "Graph"

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
	case 3:
		fmt.Println(dsMap[val])
		queue.App()
	case 4:
		fmt.Println(dsMap[val])
		binarySearchTree.App()
	case 5:
		fmt.Println(dsMap[val])
		graph.App()
	}

}
