package hashing

import (
	"fmt"
	"sort"
)

func SortArrayByOtherArrray() {
	arr1 := []int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8}
	arr2 := []int{2, 1, 8, 3}

	freqMap := make(map[int]int)
	for _, val := range arr1 {
		freqMap[val]++
	}

	for _, val := range arr2 {
		if freq, ok := freqMap[val]; ok {
			for j := 0; j < freq; j++ {
				fmt.Print(val, " ")
			}
			delete(freqMap, val)
		}
	}

	currArr := make([]int, 0, len(freqMap))

	for key, _ := range freqMap {
		currArr = append(currArr, key)
	}
	sort.Slice(currArr, func(i, j int) bool {
		return currArr[i] < currArr[j]
	})

	for _, val := range currArr {
		for j := 0; j < freqMap[val]; j++ {
			fmt.Print(val, " ")
		}
	}
	fmt.Println()
}
