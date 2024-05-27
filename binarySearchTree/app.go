package binarySearchTree

import "fmt"

func App() {

	fmt.Println("Binary search tree feature")
	bstMap := make(map[int]string)
	bstMap[0] = "Insert Search"
	for key, val := range bstMap {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case 0:
		fmt.Println(bstMap[val])
		InsertSearch()
	}

}
