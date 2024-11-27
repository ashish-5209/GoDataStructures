package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	CountProductSubarrayLessThanK() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) CountProductSubarrayLessThanK() int {
	if e.k == 0 {
		return 0
	}
	start, end, count := 0, 0, 0
	currProd := 1

	for end < len(e.arr) {
		currProd *= e.arr[end]
		for currProd >= e.k && start <= end {
			currProd /= e.arr[start]
			start++
		}
		count += (end - start + 1)
		end++
	}

	return count
}

func main() {
	obj := newElem([]int{10, 5, 2, 6}, 100)
	fmt.Println(obj.CountProductSubarrayLessThanK())
}
