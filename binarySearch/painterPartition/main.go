package main

import (
	"fmt"
)

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	PaintWall() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) PaintWall() int {
	ans := -1
	start, end := -1, 0
	for _, val := range e.arr {
		start = max(start, val)
		end += val
	}

	for start <= end {
		mid := start + (end-start)/2
		walls, count := 0, 1
		for _, val := range e.arr {
			walls += val
			if walls > mid {
				count++
				walls = val
			}
		}
		if count <= e.k {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func main() {
	arr := []int{5, 10, 30, 20, 15}
	k := 3
	obj := newElem(arr, k)
	res := obj.PaintWall()
	fmt.Println(res)
}
