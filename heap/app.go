package heap

import "fmt"

func App() {
	heap := make(map[int]string)
	heap[0] = "Insert Delete"
	for key, val := range heap {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case val:
		fmt.Println(heap[val])
		InsertDelete()
	}
}
