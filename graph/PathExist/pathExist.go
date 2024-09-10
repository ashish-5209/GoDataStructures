package main

import (
	"container/list"
	"fmt"
)

type Grid struct {
	grid [][]int
}

type GridI interface {
	CheckPathExist(int, int) bool
	IsSafe(int, int) bool
}

func newGrid() *Grid {
	return &Grid{
		grid: make([][]int, 0),
	}
}

func (g *Grid) IsSafe(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(g.grid) && y < len(g.grid[x])
}

func (g *Grid) CheckPathExist(x, y int) bool {
	q := list.New()
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, 1, -1, 0}

	q.PushBack([]int{x, y})

	visited := make([][]bool, len(g.grid))
	for i := range len(g.grid) {
		visited[i] = make([]bool, len(g.grid[i]))
	}
	visited[x][y] = true

	for q.Len() > 0 {
		curr := q.Front().Value.([]int)
		q.Remove(q.Front())

		if g.grid[curr[0]][curr[1]] == 2 {
			return true
		}

		for i := range 4 {
			nx := curr[0] + dx[i]
			ny := curr[1] + dy[i]

			if g.IsSafe(nx, ny) && g.grid[nx][ny] != 0 && !visited[nx][ny] {
				visited[nx][ny] = true
				q.PushBack([]int{nx, ny})
			}
		}
	}

	return false
}

func main() {
	g := newGrid()
	g.grid = append(g.grid, []int{3, 0, 3, 0, 0})
	g.grid = append(g.grid, []int{3, 0, 0, 0, 3})
	g.grid = append(g.grid, []int{3, 3, 3, 3, 3})
	g.grid = append(g.grid, []int{0, 2, 3, 0, 0})
	g.grid = append(g.grid, []int{3, 0, 0, 1, 3})

	sx, sy := 0, 0

	for i := range len(g.grid) {
		for j := range len(g.grid[i]) {
			if g.grid[i][j] == 1 {
				sx, sy = i, j
				break
			}
		}
	}

	val := g.CheckPathExist(sx, sy)

	fmt.Println(val)
}
