package hashing

import "fmt"

func LongestConsecutiveSubsequence() {
	arr := []int{1, 9, 3, 10, 4, 20, 2}
	boolMap := make(map[int]bool)
	ans := 0
	for _, val := range arr {
		boolMap[val] = true
	}

	for key := range boolMap {
		if _, ok := boolMap[key-1]; !ok {
			j := key
			for {
				if _, ok1 := boolMap[j]; ok1 {
					j++
				} else {
					break
				}
			}
			ans = max(ans, j-key)
		}
	}
	fmt.Println(ans)
}
