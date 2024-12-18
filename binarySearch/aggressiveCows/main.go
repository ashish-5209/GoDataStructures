package main

import (
	"fmt"
	"slices"
)

type Element struct {
	stalls   []int
	noOfCows int
}

type ElementI interface {
	PlaceAggressiveCows() int
}

func newElem(stalls []int, noOfCows int) ElementI {
	return Element{
		stalls:   stalls,
		noOfCows: noOfCows,
	}
}

func (e Element) PlaceAggressiveCows() int {
	slices.Sort(e.stalls)
	start, end := 1, e.stalls[len(e.stalls)-1]
	distance := -1
	for start <= end {
		mid := start + (end-start)/2
		count, pos := 1, e.stalls[0]
		for _, val := range e.stalls {
			if mid+pos <= val {
				count++
				pos = val
			}
		}
		if count < e.noOfCows {
			end = mid - 1
		} else {
			distance = mid
			start = mid + 1
		}
	}
	return distance
}

func main() {
	stalls := []int{2, 12, 11, 3, 26, 7}
	noOfCows := 5
	obj := newElem(stalls, noOfCows)
	res := obj.PlaceAggressiveCows()
	fmt.Println(res)
}
