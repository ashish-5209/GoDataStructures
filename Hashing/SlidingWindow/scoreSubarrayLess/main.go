package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	CountScoreSubarrayLessThanK() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) CountScoreSubarrayLessThanK() int {
	count, currSum, currLen, start, end := 0, 0, 0, 0, 0

	for end < len(e.arr) {
		currSum += e.arr[end]
		currLen++
		for {
			if currSum*currLen < e.k || start > end {
				break
			}
			currSum -= e.arr[start]
			currLen--
			start++
		}
		count += currLen
		end++
	}
	return count
}

func main() {
	obj := newElem([]int{1, 2, 3, 4, 5}, 10)
	fmt.Println(obj.CountScoreSubarrayLessThanK())
}
