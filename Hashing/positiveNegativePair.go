package hashing

import (
	"fmt"
)

func FindPositiveNegativePair() {
	arr := []int{1, 3, 6, -2, -1, -3, 2, 7}
	// Use a map to track the presence of numbers
	seen := make(map[int]bool)
	// Slice to store the result pairs
	pairs := make([][2]int, 0)

	// Iterate through the array
	for _, num := range arr {
		// If the negative counterpart exists in the map, we have a pair
		if num > 0 && seen[-num] {
			pairs = append(pairs, [2]int{-num, num})
		} else if num < 0 && seen[-num] {
			pairs = append(pairs, [2]int{num, -num})
		}
		// Mark the current number as seen
		seen[num] = true
	}

	// Print the pairs
	for _, pair := range pairs {
		fmt.Printf("%d %d ", pair[0], pair[1])
	}
	fmt.Println()
}
