package main

type Graph struct {
	edges map[int][]int
}

type GraphI interface{}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[int][]int),
	}
}

func (g *Graph) AddVertex(src, dest int) {
	g.edges[src] = append(g.edges[src], dest)
}

func (g Graph) IsSafe(src, c, V int, color []int) bool {
	for _, val := range g.edges[src] {
		if color[val] == c {
			return false
		}
	}
	return true
}

func (g Graph) GraphColoringUtil(m, v, V int, color []int) bool {
	if v == V {
		return true
	}
	for c := 0; c <= m; c++ {
		if g.IsSafe(v, c, V, color) {
			color[v] = c
			if g.GraphColoringUtil(m, v+1, V, color) {
				return true
			}
			color[v] = 0
		}
	}
	return false
}

func (g Graph) GraphColoring(k, V int) bool {
	color := make([]int, V)
	return g.GraphColoringUtil(k, 0, V, color)
}

func main() {
	graph := NewGraph()
	graph.AddVertex(0, 1)
	graph.AddVertex(0, 2)
	graph.AddVertex(1, 2)
	graph.AddVertex(2, 3)
	graph.AddVertex(3, 0)
}
