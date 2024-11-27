package main

import "fmt"

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	TotalSubrrayWithKDIfferentIntegersUtil() int
	TotalSubrrayWithKDIfferentIntegers() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (e Element) TotalSubrrayWithKDIfferentIntegersUtil() (total int) {
	start, end, total, count := 0, 0, 0, 0
	arrMap := make(map[int]int)

	for end < len(e.arr) {
		arrMap[e.arr[end]]++
		if arrMap[e.arr[end]] == 1 {
			count++
		}
		for count == e.k && start <= end {
			total += len(e.arr) - end
			arrMap[e.arr[start]]--
			if arrMap[e.arr[start]] == 0 {
				count--
			}
			start++
		}
		end++
	}
	return total
}

func (e Element) TotalSubrrayWithKDIfferentIntegers() (total int) {
	total1 := e.TotalSubrrayWithKDIfferentIntegersUtil()
	e.k++
	total2 := e.TotalSubrrayWithKDIfferentIntegersUtil()
	return total1 - total2
}

func main() {
	obj := newElem([]int{1, 2, 1, 2, 3}, 2)
	fmt.Println(obj.TotalSubrrayWithKDIfferentIntegers())
}
