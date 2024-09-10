package main

import "fmt"

type Grid struct {
	grid [][]int
}

type GridI interface {
	Island() int
	DFS(int, int)
}

func newGrid() *Grid {
	return &Grid{
		grid: make([][]int, 0),
	}
}

func isSafe(x, y, n, m int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func (g *Grid) DFS(x, y int) {
	dx := []int{1, 1, 1, -1, -1, -1, 0, 0}
	dy := []int{1, -1, 0, 1, -1, 0, 1, -1}

	for i := range 8 {
		if isSafe(x+dx[i], y+dy[i], len(g.grid), len(g.grid[0])) && g.grid[x+dx[i]][y+dy[i]] == 1 {
			g.grid[x+dx[i]][y+dy[i]] = 2
			g.DFS(x+dx[i], y+dy[i])
		}
	}
}

func (g *Grid) Island() int {
	cnt := 0
	for i := range len(g.grid) {
		for j := range len(g.grid[0]) {
			if g.grid[i][j] == 1 {
				g.grid[i][j] = 2
				g.DFS(i, j)
				cnt += 1
			}
		}
	}

	return cnt
}

func main() {
	g := newGrid()
	g.grid = append(g.grid, []int{0, 1, 1, 1, 0, 0, 0})
	g.grid = append(g.grid, []int{0, 0, 1, 1, 0, 1, 0})
	// g.grid = append(g.grid, []int{1, 1})
	// g.grid = append(g.grid, []int{1, 0})
	val := g.Island()
	fmt.Println(val)
}
