package stringDS

import (
	"fmt"
	"strconv"
)

func FindNthPrimeDigitNumber() {
	n := 21
	primes := []int{2, 3, 5, 7}
	queue := []string{"2", "3", "5", "7"}

	for i := 1; i < n; i++ {
		curr := queue[0]
		queue = queue[1:]
		for _, prime := range primes {
			newNumber := curr + strconv.Itoa(prime)
			queue = append(queue, newNumber)
		}
	}
	fmt.Println(strconv.Atoi(queue[0]))
}
