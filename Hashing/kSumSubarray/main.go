package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	CountKSumSubarray() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) CountKSumSubarray() int {
	count, currSum := 0, 0
	prefixSumMap := make(map[int]int)

	for _, val := range e.arr {
		currSum += val
		if currSum == e.k {
			count++
		}
		if freq, ok := prefixSumMap[currSum-e.k]; ok {
			count += freq
		}
		prefixSumMap[currSum]++
	}

	return count
}

func main() {
	obj := newElem([]int{10, 2, -2, -20, 10}, -10)
	fmt.Println(obj.CountKSumSubarray())
}
