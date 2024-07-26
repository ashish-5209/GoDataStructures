package hashing

import "fmt"

func PairSumDistinct() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	elem := 14
	arrMap := make(map[int]int)

	for idx, val := range arr {
		curr := elem - val
		if index, ok := arrMap[curr]; ok {
			fmt.Printf("Total %d available at index %d, %d with value %d + %d", elem, index, idx, arr[index], arr[idx])
			fmt.Println()
		}
		arrMap[val] = idx
	}
}
