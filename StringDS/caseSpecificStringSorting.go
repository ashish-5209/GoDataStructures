package stringDS

import (
	"fmt"
	"sort"
)

func CaseSpecificStringSorting() {
	s := "defRTSersUXI"
	var arr1, arr2 []rune
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			arr1 = append(arr1, ch)
		} else {
			arr2 = append(arr2, ch)
		}
	}

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	var result []rune
	i, j := 0, 0
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	fmt.Println(string(result))
}
