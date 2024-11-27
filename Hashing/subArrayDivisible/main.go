package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	CountSubarrayDivisibleByK() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) CountSubarrayDivisibleByK() int {
	count, currSum, rem := 0, 0, 0
	prefixRemMap := make(map[int]int)
	prefixRemMap[0] = 1

	for _, val := range e.arr {
		currSum += val
		rem = (currSum%e.k + e.k) % e.k
		count += prefixRemMap[rem]
		prefixRemMap[rem]++
	}
	return count
}

func main() {
	obj := newElem([]int{4, 5, 0, -2, -3, 1}, 5)
	fmt.Println(obj.CountSubarrayDivisibleByK())
}
