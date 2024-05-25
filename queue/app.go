package queue

import "fmt"

func App() {
	queue := make(map[int]string)
	queue[0] = "Insert Delete"
	for key, val := range queue {
		fmt.Println(key, "			", val)
	}

	var val int
	fmt.Scanln(&val)
	switch val {
	case 0:
		fmt.Println(queue[val])
		InsertDelete()
	}
}
