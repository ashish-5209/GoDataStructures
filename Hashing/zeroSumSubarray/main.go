package main

import "fmt"

type Element struct {
	arr []int
}

type ElementI interface {
	CountZeroSumSubarray() int
}

func newElem(arr []int) ElementI {
	return Element{
		arr: arr,
	}
}

func (e Element) CountZeroSumSubarray() int {
	count := 0
	prefixSumMap := make(map[int]int)
	prefixSumMap[0]++
	currSum := 0

	for _, val := range e.arr {
		currSum += val
		if freq, ok := prefixSumMap[currSum]; ok {
			count += freq
		}
		prefixSumMap[currSum]++
	}

	return count
}

func main() {
	obj := newElem([]int{6, -1, -3, 4, -2, 2, 4, 6, -12, -7})
	fmt.Println(obj.CountZeroSumSubarray())
}
