package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	GetLongetstSubararyWithKFreq() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) GetLongetstSubararyWithKFreq() int {
	start, end, count := 0, 0, 0
	arrMap := make(map[int]int)

	for end < len(e.arr) {
		arrMap[e.arr[end]]++
		for arrMap[e.arr[end]] > e.k && start <= end {
			arrMap[e.arr[start]]--
			start++
		}
		count = max(count, end-start+1)
		end++
	}
	return count
}

func main() {
	obj := newElem([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)
	fmt.Println(obj.GetLongetstSubararyWithKFreq())
}
