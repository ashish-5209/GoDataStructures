package main

import "fmt"

type LeftIndex struct {
	arr []int
}

type LeftIndexI interface {
	FindLeftIndex(int) int
}

func newLeftIndex() *LeftIndex {
	return &LeftIndex{
		arr: make([]int, 0),
	}
}

func (l LeftIndex) FindLeftIndex(k int) int {
	start, end, mid, res := 0, len(l.arr)-1, 0, -1

	for start <= end {
		mid = start + (end-start)/2
		if l.arr[mid] == k {
			res = mid
			end = mid - 1
		} else if l.arr[mid] > k {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return res
}

func main() {
	l := newLeftIndex()
	l.arr = []int{10, 20, 20, 20, 20, 20, 20}
	val := l.FindLeftIndex(20)
	fmt.Println(val)
}
