package main

import "fmt"

type Element struct {
	piles []int
	hours int
}

type ElementI interface {
	MinimumSpeedOfEatingBnanas() int
}

func newElem(piles []int, hours int) ElementI {
	return Element{
		piles: piles,
		hours: hours,
	}
}

func (e Element) MinimumSpeedOfEatingBnanas() int {
	speed := -1
	start, end, sum := 0, e.piles[0], 0
	for _, val := range e.piles {
		sum += val
		end = max(end, val)
	}
	start = sum / e.hours
	if start == 0 {
		start = 1
	}

	for start <= end {
		mid := start + (end-start)/2
		time := 0
		for _, val := range e.piles {
			time += val / mid
			if val%mid != 0 {
				time++
			}
		}
		if time > e.hours {
			start = mid + 1
		} else {
			speed = mid
			end = mid - 1
		}
	}
	return speed
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	hours := 5

	obj := newElem(piles, hours)
	res := obj.MinimumSpeedOfEatingBnanas()
	fmt.Println(res)
}
