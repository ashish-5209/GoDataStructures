package stringDS

import "fmt"

func CountNoOfOneCombination() {
	s := "01101"
	count := 0

	for _, ch := range s {
		if ch == '1' {
			count++
		}
	}

	totalCombination := (count * (count - 1)) / 2

	fmt.Println(totalCombination)
}
