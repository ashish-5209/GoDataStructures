package linkedLst

import "fmt"

func App() {
	fmt.Println("LinkedList feature")
	llMap := make(map[int]string)
	llMap[0] = "Insert Delete"
	llMap[1] = "LRU Cache"
	for key, val := range llMap {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case 0:
		fmt.Println(llMap[val])
		InsertDelete()
	case 1:
		fmt.Println(llMap[val])
		LRUCacheDriver()
	}
}
