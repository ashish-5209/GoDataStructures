package hashing

import "fmt"

func CountNonRepeatingElem() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 7}
	arrMap := make(map[int]int)

	for _, val := range arr {
		arrMap[val]++
	}

	count := 0
	for key, val := range arrMap {
		if val == 1 {
			fmt.Print(key, " ")
			count++
		}
	}
	fmt.Println()
	fmt.Println("Total no of non repeating element is", count)
}
