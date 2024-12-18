package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	AllocateBook() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) AllocateBook() int {
	if len(e.arr) < e.k {
		return -1
	}
	ans := -1
	start, end := -1, 0
	for _, val := range e.arr {
		start = max(start, val)
		end += val
	}

	for start <= end {
		mid := start + (end-start)/2
		pages, count := 0, 1
		for _, val := range e.arr {
			pages += val
			if pages > mid {
				count++
				pages = val
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
	arr := []int{12, 34, 67, 90}
	k := 2
	obj := newElem(arr, k)
	res := obj.AllocateBook()
	fmt.Println(res)
}
