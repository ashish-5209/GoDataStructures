package graph

func (g *Graph) DetectCycleG1(s int, visited, rest []bool) bool {
	visited[s] = true
	rest[s] = true

	for _, edge := range g.edges[s] {
		if !visited[edge] && g.DetectCycleG1(edge, visited, rest) {
			return true
		} else if rest[edge] {
			return true
		}
	}
	rest[s] = false
	return false
}
