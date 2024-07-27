package hashing

import "fmt"

func LongestSubarraySum() {
	arr := []int{10, 5, 2, 7, 1, 9}
	k := 15
	idxMap := make(map[int]int)
	res := 0
	prefixSum := 0

	for idx, val := range arr {
		prefixSum += val
		if prefixSum == k {
			res = idx + 1
		}
		if _, ok := idxMap[prefixSum]; !ok {
			idxMap[prefixSum] = idx
		}
		if key, ok := idxMap[prefixSum-k]; ok {
			res = max(res, idx-key)
		}
	}
	fmt.Println(res)
}
