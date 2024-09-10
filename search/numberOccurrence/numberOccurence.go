package main

import "fmt"

type Elem struct {
	arr []int
}

type ElemI interface {
	FindOccurence(int) int
}

func newElem() *Elem {
	return &Elem{
		arr: make([]int, 0),
	}
}

func (e Elem) FindOccurence(k int) int {
	eMap := make(map[int]int)
	for _, val := range e.arr {
		eMap[val]++
	}

	for key, val := range eMap {
		if val > len(e.arr)/k {
			return key
		}
	}
	return -1
}

func main() {
	e := newElem()
	e.arr = []int{2, 3, 3, 2}
	val := e.FindOccurence(3)
	fmt.Println(val)
}
