package hashing

import "fmt"

func FirstRepeatingElement() {
	arr := []int{1, 5, 3, 4, 3, 5, 6}
	freqMap := make(map[int]int)
	idx := -1

	for _, val := range arr {
		freqMap[val]++
	}

	for _, val := range arr {
		if freqMap[val] > 1 {
			idx = freqMap[val]
			break
		}
	}

	fmt.Println(idx)
}
