package main

import "fmt"

type Element struct {
	arr    []int
	target int
}

type ElementI interface {
	FindElementIndex() int
}

func newElem(arr []int, target int) ElementI {
	return Element{
		arr:    arr,
		target: target,
	}
}

func (e Element) FindElementIndex() int {
	start, end := 0, len(e.arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if e.arr[mid] == e.target {
			return mid
		} else if e.arr[mid] > e.target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	obj := newElem(arr, target)
	fmt.Println(obj.FindElementIndex())
}
