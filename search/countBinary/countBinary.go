package main

import "fmt"

type Elem struct {
	arr []int
}

type ElemI interface {
	CountBinary() int
}

func newElem() *Elem {
	return &Elem{
		arr: make([]int, 0),
	}
}

func (e Elem) CountBinary() int {
	if e.arr[0] == 0 {
		return 0
	}
	start, end, mid, res := 0, len(e.arr)-1, 0, 0

	for start <= end {
		mid = start + (end-start)/2
		if e.arr[mid] == 1 {
			res = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return res + 1
}

func main() {
	e := newElem()
	e.arr = []int{1, 1, 1, 1, 1, 0, 0, 0}
	val := e.CountBinary()
	fmt.Println(val)
}
