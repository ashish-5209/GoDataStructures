package main

import "fmt"

type Element struct {
	arr    []int
	target int
}

type ElementI interface {
	GetMinimumSizeSubarraySum() int
}

func newElem(arr []int, target int) ElementI {
	return Element{
		arr:    arr,
		target: target,
	}
}

func (e Element) GetMinimumSizeSubarraySum() int {
	prefixSum := 0
	start, end, count := 0, 0, len(e.arr)+1

	for end < len(e.arr) {
		prefixSum += e.arr[end]
		for prefixSum >= e.target && start <= end {
			count = min(count, end-start+1)
			prefixSum -= e.arr[start]
			start++
		}
		end++
	}
	if count == len(e.arr)+1 {
		return 0
	}
	return count
}

func main() {
	obj := newElem([]int{2, 3, 1, 2, 4, 3}, 7)
	fmt.Println(obj.GetMinimumSizeSubarraySum())
}
