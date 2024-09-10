package main

import (
	"container/list"
	"fmt"
)

type Grid struct {
	grid [][]int
}

type GridI interface {
	RottenOrange() int
}

func newGrid() *Grid {
	return &Grid{
		grid: make([][]int, 0),
	}
}

func (g *Grid) RottenOrange() int {
	q := list.New()

	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid[i]); j++ {
			if g.grid[i][j] == 2 {
				q.PushBack([]int{i, j})
			}
		}
	}

	totalTime := 0
	for q.Len() > 0 {
		l := q.Len()
		flag := false

		for l > 0 {
			l--
			curr := q.Front().Value.([]int)
			q.Remove(q.Front())
			x, y := curr[0], curr[1]
			g.grid[x][y] = 0

			if (x-1) >= 0 && g.grid[x-1][y] == 1 {
				q.PushBack([]int{x - 1, y})
				flag = true
			}
			if (y-1) >= 0 && g.grid[x][y-1] == 1 {
				q.PushBack([]int{x, y - 1})
				flag = true
			}
			if (x+1) < len(g.grid) && g.grid[x+1][y] == 1 {
				q.PushBack([]int{x + 1, y})
				flag = true
			}

			if (y+1) < len(g.grid[x]) && g.grid[x][y+1] == 1 {
				q.PushBack([]int{x, y + 1})
				flag = true
			}
		}
		if flag {
			totalTime++
		}
	}

	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid[i]); j++ {
			if g.grid[i][j] == 1 {
				return -1
			}
		}
	}

	return totalTime
}

func main() {
	// grid := [][]int {{0,1,2},{0,1,2},{2,1,1}}
	grid := newGrid()
	grid.grid = append(grid.grid, []int{0, 1, 2})
	grid.grid = append(grid.grid, []int{0, 1, 2})
	grid.grid = append(grid.grid, []int{2, 1, 1})

	totalTime := grid.RottenOrange()
	fmt.Println(totalTime)
}
