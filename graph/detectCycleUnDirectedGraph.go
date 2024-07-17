package graph

func (g *Graph) DetectCycleUG(s int, parent int, visited []bool) bool {
	if !g.vertices[s] {
		return true
	}
	visited[s] = true
	for _, edge := range g.edges[s] {
		if !visited[edge] {
			if g.DetectCycleUG(edge, s, visited) {
				return true
			}
		}
		if edge != parent {
			return true
		}
	}
	return false
}
