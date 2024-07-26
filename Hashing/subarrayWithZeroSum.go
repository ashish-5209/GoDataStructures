package hashing

import "fmt"

func SubArrayWithZeroSum() {
	arr := []int{4, 2, -3, 1, 6}
	prefixMap := make(map[int]bool)
	prefixSum := 0
	prefixMap[prefixSum] = true

	for _, val := range arr {
		prefixSum += val
		if prefixMap[prefixSum] {
			fmt.Println("YES")
			return
		}
		prefixMap[prefixSum] = true
	}
	fmt.Println("No")
}
