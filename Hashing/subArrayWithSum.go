package hashing

import "fmt"

func SubArrayWithGivenSum() {
	arr := []int{10, 2, -2, -20, 10}
	prefixMap := make(map[int]int)
	currSum := 0
	count := 0
	sum := -10

	for _, val := range arr {
		currSum += val
		if currSum == sum {
			count++
		}
		if freq, ok := prefixMap[currSum-sum]; ok {
			count += freq
		}
		prefixMap[currSum]++
	}

	fmt.Println(count)
}
