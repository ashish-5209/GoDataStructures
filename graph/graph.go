package graph

import "fmt"

type Graph struct {
	vertices map[int]bool
	edges    map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int]bool),
		edges:    make(map[int][]int),
	}
}

func (g *Graph) AddVertex(v int) {
	if !g.vertices[v] {
		g.vertices[v] = true
		g.edges[v] = []int{}
	}
}

func (g *Graph) AddEdges(v1, v2 int) {
	if !g.vertices[v1] || !g.vertices[v2] {
		fmt.Println("Error: Both vertices must exist in the graph.")
		return
	}

	g.edges[v1] = append(g.edges[v1], v2)
	g.edges[v2] = append(g.edges[v2], v1)
}

func (g *Graph) AddEdgesDirctional(v1, v2 int) {
	if !g.vertices[v1] || !g.vertices[v2] {
		fmt.Println("Error: Both vertices must exist in the graph.")
		return
	}

	g.edges[v1] = append(g.edges[v1], v2)
}

func (g *Graph) Display() {
	for v, neighbour := range g.edges {
		fmt.Printf("%d -> %v", v, neighbour)
		fmt.Println()
	}
}

func main() {
	graph := NewGraph()
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdges(1, 2)
	graph.AddEdges(1, 3)
	graph.AddEdges(2, 4)
	graph.AddEdges(3, 4)

	fmt.Println(len(graph.vertices))

	graph.Display()
	graph.BFS()
	fmt.Println()

	var visited []bool
	x := 0
	for x <= len(graph.edges) {
		visited = append(visited, false)
		x += 1
	}

	graph.DFS(1, visited)
	fmt.Println()
	graph.ShortestPathUnweightedGraph(1)

	x = 0
	for x <= len(graph.edges) {
		visited[x] = false
		x += 1
	}

	fmt.Println(graph.DetectCycleUG(1, -1, visited))

	directionalGraph := NewGraph()
	directionalGraph.AddVertex(1)
	directionalGraph.AddVertex(2)
	directionalGraph.AddVertex(3)
	directionalGraph.AddVertex(4)
	directionalGraph.AddVertex(5)
	directionalGraph.AddVertex(6)

	directionalGraph.AddEdgesDirctional(1, 2)
	directionalGraph.AddEdgesDirctional(3, 2)
	directionalGraph.AddEdgesDirctional(3, 4)
	directionalGraph.AddEdgesDirctional(4, 5)
	directionalGraph.AddEdgesDirctional(5, 6)
	directionalGraph.AddEdgesDirctional(6, 4)

	var visitedDirectional, rest []bool
	x = 0
	for x <= len(directionalGraph.vertices) {
		visitedDirectional = append(visitedDirectional, false)
		rest = append(rest, false)
	}
	res := false
	for key := range directionalGraph.vertices {
		if !visitedDirectional[key] {
			if directionalGraph.DetectCycleG1(key, visited, rest) {
				res = true
			} else {
				res = false
				break
			}
		}
	}
	fmt.Println(res)
}
