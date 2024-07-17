package graph

import "fmt"

func (n *Graph) DFS(s int, visited []bool) {
	if !n.vertices[s] {
		return
	}
	visited[s] = true
	fmt.Print(s, " ")

	for _, edge := range n.edges[s] {
		if !visited[edge] {
			n.DFS(edge, visited)
		}
	}
}
