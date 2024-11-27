package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	CountSubarrayOfMaxElemKTimes() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) CountSubarrayOfMaxElemKTimes() (total int) {
	maxVal := 0
	for _, val := range e.arr {
		maxVal = max(maxVal, val)
	}
	start, end, count := 0, 0, 0
	for end < len(e.arr) {
		if e.arr[end] == maxVal {
			count++
		}
		for count == e.k && start <= end {
			total += len(e.arr) - end
			if e.arr[start] == maxVal {
				count--
			}
			start++
		}

		end++
	}
	return total
}

func main() {
	obj := newElem([]int{1, 3, 2, 3, 3}, 2)
	fmt.Println(obj.CountSubarrayOfMaxElemKTimes())
}
