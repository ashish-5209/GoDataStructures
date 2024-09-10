package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	edges map[int][]int
}

func newGraph() *Graph {
	return &Graph{
		edges: make(map[int][]int),
	}
}

type GraphI interface {
	AddEdges(int, int)
	NodeLevel(int) int
}

func (g *Graph) AddEdges(u, v int) {
	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}

func (g *Graph) NodeLevel(s int) int {
	q := list.New()
	level := 0
	visited := make([]bool, len(g.edges))
	visited[0] = true
	q.PushBack(0)

	for q.Len() > 0 {
		l := q.Len()
		for l > 0 {
			l--
			curr := q.Front().Value.(int)
			q.Remove(q.Front())
			if curr == s {
				return level
			}
			for _, edge := range g.edges[curr] {
				if !visited[edge] {
					q.PushBack(edge)
					visited[edge] = true
				}
			}
		}
		level++
	}
	return -1
}

func main() {
	g := newGraph()
	g.AddEdges(0, 1)
	g.AddEdges(0, 2)
	g.AddEdges(2, 5)
	g.AddEdges(1, 3)
	g.AddEdges(1, 4)

	val := g.NodeLevel(4)
	fmt.Println(val)
}
