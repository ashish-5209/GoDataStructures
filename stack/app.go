package stack

import "fmt"

func App() {
	stack := make(map[int]string)
	stack[0] = "Insert Delete"

	for key, val := range stack {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)
	switch val {
	case 0:
		fmt.Println(stack[val])
		InsertDelete()
	}
}
