package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 2, 2}
	res := -1

	totalSum := 0
	for i := 0; i < len(arr); i++ {
		totalSum += arr[i]
	}
	lSum := 0
	for i := 0; i < len(arr); i++ {
		totalSum -= arr[i]
		if lSum == totalSum {
			res = i
			fmt.Println(lSum)
			break
		}
		lSum += arr[i]
	}

	fmt.Println(res)
}
