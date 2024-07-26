package hashing

import "fmt"

func UnionOfArray() {
	arr1 := []int{85, 25, 1, 32, 54, 6}
	arr2 := []int{85, 2}

	freqMap := make(map[int]bool)

	for _, val := range arr1 {
		freqMap[val] = true
	}

	for _, val := range arr2 {
		if _, ok := freqMap[val]; !ok {
			freqMap[val] = true
		}
	}

	fmt.Println(len(freqMap))
}
