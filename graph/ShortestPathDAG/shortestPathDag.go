package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	vertices map[int]bool
	edges    map[int][][]int
}

type ShortestPathI interface {
	AddVertex(int)
	AddEdges(int, int, int)
	TopologicalSort(int, *list.List, []bool)
	FindShortestPathDAG()
}

func newGraph() *Graph {
	return &Graph{
		vertices: make(map[int]bool),
		edges:    make(map[int][][]int),
	}
}

func (g *Graph) AddVertex(v int) {
	if !g.vertices[v] {
		g.vertices[v] = true
		g.edges[v] = [][]int{}
	}
}

func (g *Graph) AddEdges(u, v, w int) {
	if !g.vertices[u] && !g.vertices[v] {
		fmt.Println("vertex are not added")
		return
	}

	g.edges[u] = append(g.edges[u], []int{v, w})
}

func (g *Graph) TopologicalSort(s int, stack *list.List, visited []bool) {
	visited[s] = true
	for _, edge := range g.edges[s] {
		if !visited[edge[0]] {
			g.TopologicalSort(edge[0], stack, visited)
		}
	}
	stack.PushFront(s)
}

func (g *Graph) FindShortestPathDAG() {
	stack := list.New()
	visited := make([]bool, len(g.vertices))

	for x := 0; x < len(g.vertices); x++ {
		if !visited[x] {
			g.TopologicalSort(x, stack, visited)
		}
	}

	var dist []int
	for x := 0; x < len(g.vertices); x++ {
		dist = append(dist, 1000000)
	}

	dist[1] = 0

	for stack.Len() > 0 {
		curr := stack.Front()
		stack.Remove(curr)
		curr_node := curr.Value.(int)
		for _, val := range g.edges[curr_node] {
			fmt.Println("Val", curr_node, val, dist[val[0]], dist[curr_node]+val[1])
			if dist[val[0]] > dist[curr_node]+val[1] {
				dist[val[0]] = dist[curr_node] + val[1]
			}
		}
	}

	fmt.Println(dist)
}

func main() {
	graph := newGraph()
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)

	graph.AddEdges(0, 1, 5)
	graph.AddEdges(0, 2, 3)
	graph.AddEdges(1, 3, 6)
	graph.AddEdges(1, 2, 2)
	graph.AddEdges(2, 4, 4)
	graph.AddEdges(2, 5, 2)
	graph.AddEdges(2, 3, 7)
	graph.AddEdges(3, 4, -1)
	graph.AddEdges(4, 5, -2)

	fmt.Println(graph)

	graph.FindShortestPathDAG()
}
