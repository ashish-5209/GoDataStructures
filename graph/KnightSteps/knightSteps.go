package main

import (
	"container/list"
	"fmt"
)

type Cell struct {
	x, y, dist int
}

// type Grid struct {
// 	gird [][]int
// }

// type GridI interface {
// 	MinStepsKnight([]int, []int, int) int
// }

// func newGrid() *Grid {
// 	return &Grid{
// 		gird: make([][]int, 0),
// 	}
// }

func isInside(x, y, n int) bool {
	if x > 0 && x <= n && y > 0 && y <= n {
		return true
	}
	return false
}

func MinStepsKnight(src, dest []int, n int) int {
	dx := []int{2, 2, -2, -2, 1, 1, -1, -1}
	dy := []int{1, -1, 1, -1, 2, -2, 2, -2}

	q := list.New()

	q.PushBack(Cell{
		x:    src[0],
		y:    src[1],
		dist: 0,
	})

	visited := make([][]bool, n+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
	}

	visited[src[0]][src[1]] = true

	for q.Len() > 0 {
		curr := q.Front().Value.(Cell)
		q.Remove(q.Front())

		if curr.x == dest[0] && curr.y == dest[1] {
			return curr.dist
		}

		for i := range 8 {
			x := curr.x + dx[i]
			y := curr.y + dy[i]

			if isInside(x, y, n) && !visited[x][y] {
				visited[x][y] = true
				q.PushBack(Cell{
					x:    x,
					y:    y,
					dist: curr.dist + 1,
				})
			}
		}
	}
	return -1
}

func main() {
	n := 6
	val := MinStepsKnight([]int{4, 5}, []int{1, 1}, n)
	fmt.Println(val)
}
