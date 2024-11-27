package main

import (
	"fmt"
	"sort"
)

type Element struct {
	arr    []int
	target int
}

type ElementI interface {
	CombinationSum()
	CombinationSumHelper(*[][]int, int, []int, int)
}

func newElem(arr []int, target int) ElementI {
	return Element{
		arr:    arr,
		target: target,
	}
}

func (e Element) CombinationSumHelper(res *[][]int, currSum int, targetArr []int, start int) {
	if currSum == e.target {
		*res = append(*res, append([]int{}, targetArr...))
		return
	}

	for i := start; i < len(e.arr); i++ {
		// Choose element e.arr[i] if it doesn’t exceed target
		if currSum+e.arr[i] <= e.target {
			targetArr = append(targetArr, e.arr[i])
			e.CombinationSumHelper(res, currSum+e.arr[i], targetArr, i)
			// Backtrack
			targetArr = targetArr[:len(targetArr)-1]
		}
	}
}

func (e Element) CombinationSum() {
	var res [][]int
	var targetArr []int

	e.CombinationSumHelper(&res, 0, targetArr, 0)

	fmt.Println(res)
}

func main() {
	arr := []int{7, 2, 6, 5}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	obj := newElem(arr, 16)
	obj.CombinationSum()
}
