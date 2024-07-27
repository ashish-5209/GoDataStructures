package graph

import "fmt"

func App() {
	graph := make(map[int]string)
	graph[0] = "Insert Print"
	graph[1] = "Graph Logic"
	for key, val := range graph {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case 0:
		fmt.Println(graph[val])
		InsertPrint()
	case 1:
		fmt.Println(graph[val])
		GraphDS()
	}
}
