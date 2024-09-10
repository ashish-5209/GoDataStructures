package main

import "fmt"

type Floor struct {
	Arr []int
}

type FloorI interface {
	FindFloor(int) int
}

func newFloor() *Floor {
	return &Floor{
		Arr: make([]int, 0),
	}
}

func (f Floor) FindFloor(val int) int {
	start, end, mid := 0, len(f.Arr)-1, 0
	res := -1

	for start <= end {
		mid = start + (end-start)/2
		if f.Arr[mid] == val {
			return mid
		}
		if f.Arr[mid] < val {
			start = mid + 1
		} else {
			end = mid - 1
			res = end
		}
	}
	return res
}

func main() {
	floor := newFloor()
	floor.Arr = []int{1, 2, 8, 10, 11, 12, 19}
	fmt.Println(floor.FindFloor(0))
}
