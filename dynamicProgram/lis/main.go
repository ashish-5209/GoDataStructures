package main

import "fmt"

type Element struct {
	arr []int
}

type ElementI interface {
	LISLogSoln() int
	LISSoln() int
}

func newElem(arr []int) ElementI {
	return Element{
		arr: arr,
	}
}

func Ceil(tail []int, start, end, x int) int {
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		if tail[mid] >= x {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func (e Element) LISLogSoln() int {
	var tail []int
	length := 1
	tail = append(tail, e.arr[0])
	for idx, val := range e.arr {
		if idx == 0 {
			continue
		}
		if val > tail[length-1] {
			length++
			tail = append(tail, val)
		} else {
			c := Ceil(tail, 0, length-1, val)
			tail[c] = val
		}
	}
	return length
}

func (e Element) LISSoln() int {
	lis := make([]int, len(e.arr))
	maxLis := 1
	// Set default value to 1
	for i := range lis {
		lis[i] = 1
	}

	for i := 0; i < len(e.arr); i++ {
		for j := 0; j < i; j++ {
			if e.arr[i] > e.arr[j] {
				lis[i] = max(lis[i], lis[j]+1)
				maxLis = max(maxLis, lis[i])
			}
		}
	}
	return maxLis
}

func main() {
	arr := []int{3, 10, 2, 1, 20}
	obj := newElem(arr)
	fmt.Println(obj.LISLogSoln())
	fmt.Println(obj.LISSoln())
}
