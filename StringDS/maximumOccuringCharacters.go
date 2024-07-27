package stringDS

import "fmt"

func MaximumOccurrenceCharacter() {
	s := "output"
	arr := make([]int, 26)
	maxCount := 0
	var maxChar rune

	for _, ch := range s {
		index := ch - 'a'
		arr[index]++
		if arr[index] > maxCount {
			maxCount = arr[index]
			maxChar = ch
		} else if arr[index] == maxCount && ch < maxChar {
			maxChar = ch
		}
	}
	fmt.Println(string(maxChar))
}
