package graph

import (
	"container/list"
	"fmt"
)

func (n *Graph) BFS() {
	l := list.New()
	l.PushBack(1)
	var visited []bool
	x := 0
	for x <= len(n.edges) {
		visited = append(visited, false)
		x += 1
	}
	visited[1] = true

	for l.Len() > 0 {
		curr := l.Front().Value.(int)
		l.Remove(l.Front())

		fmt.Print(curr, " ")

		for _, edge := range n.edges[curr] {
			if !visited[edge] {
				l.PushBack(edge)
				visited[edge] = true
			}
		}
	}
}
