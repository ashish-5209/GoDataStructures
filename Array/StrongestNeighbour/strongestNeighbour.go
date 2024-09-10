package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 2, 3, 4, 5}
	var result []int
	for i := 1; i < len(arr); i++ {
		result = append(result, max(arr[i], arr[i-1]))
	}

	fmt.Println(result)
}
