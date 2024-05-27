package graph

import "fmt"

func App() {
	graph := make(map[int]string)
	graph[0] = "Insert Print"
	for key, val := range graph {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case val:
		fmt.Println(graph[val])
		InsertPrint()
	}
}
