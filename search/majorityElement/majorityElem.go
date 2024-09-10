package main

import "fmt"

type MajorElem struct {
	arr []int
}

type MajorElemI interface {
	FindMajorityElement() int
}

func newMajorElem() *MajorElem {
	return &MajorElem{
		arr: make([]int, 0),
	}
}

func (m MajorElem) FindMajorityElement() int {
	count, candidate := 0, -1

	for _, val := range m.arr {
		if count == 0 {
			candidate = val
			count++
		} else {
			if candidate == val {
				count++
			} else {
				count--
			}
		}
	}

	count = 0
	for _, val := range m.arr {
		if candidate == val {
			count++
		}
	}

	if count > len(m.arr)/2 {
		return candidate
	}
	return -1
}

func main() {
	m := newMajorElem()
	m.arr = []int{3, 1, 3, 3, 2}
	val := m.FindMajorityElement()
	fmt.Println(val)
}
