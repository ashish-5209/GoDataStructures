package main

import "fmt"

type Search struct {
	data []int
}

type SearchI interface {
	BinarySearch(int) int
}

func newSearch() *Search {
	return &Search{
		data: make([]int, 0),
	}
}

func (s Search) BinarySearch(val int) int {
	start, end, mid := 0, len(s.data), 0
	for start <= end {
		mid = start + (end-start)/2
		if s.data[mid] == val {
			return mid
		}
		if s.data[mid] < val {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	search := newSearch()
	search.data = []int{1, 3, 4, 5, 6}

	fmt.Println(search.BinarySearch(2))
}
