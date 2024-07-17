package graph

import (
	"container/list"
	"fmt"
)

func (g *Graph) ShortestPathUnweightedGraph(s int) {
	if !g.vertices[s] {
		return
	}

	var visited []bool
	x := 0
	for x <= len(g.vertices) {
		visited = append(visited, false)
		x += 1
	}
	var dist []int
	x = 0
	for x <= len(g.vertices) {
		dist = append(dist, 0)
		x += 1
	}

	l := list.New()
	l.PushBack(s)
	visited[s] = true
	for l.Len() > 0 {
		curr := l.Front().Value.(int)
		l.Remove(l.Front())

		for _, edge := range g.edges[curr] {
			if !visited[edge] {
				dist[edge] = dist[curr] + 1
				visited[edge] = true
				l.PushBack(edge)
			}
		}
	}

	fmt.Println(dist)
}
