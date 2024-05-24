package main

import (
	"fmt"
	linkedList "gods/LinkedList"
	"gods/heap"
)

func main() {
	dsMap := make(map[int]string)
	dsMap[0] = "Linked List"
	dsMap[1] = "Heap"

	for key, val := range dsMap {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case 0:
		fmt.Println(dsMap[0])
		linkedList.App()
	case 1:
		fmt.Println(dsMap[1])
		heap.App()
	}
}
