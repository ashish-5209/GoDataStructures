package graph

import "fmt"

// Graphs represents an adjacency list graphs
type Graphs struct {
	vertices []*Vertex
}

// Vertex represents a graphs vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a vertex to the graphs
func (g *Graphs) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// AddEdge adds and edge to the graphs
func (g *Graphs) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// Check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// Add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex return a pointer to the vertex with a key integer
func (g *Graphs) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent List for each vertex of the origin
func (g *Graphs) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.key)
		}
	}
	fmt.Println()
}

func InsertPrint() {
	test := &Graphs{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 2)
	test.Print()
}
