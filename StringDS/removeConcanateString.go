package stringDS

import "fmt"

func RemoveAndConcanateString() {
	s1 := "abcs"
	s2 := "cxzca"
	res := ""
	var arr [26]int

	for _, ch := range s1 {
		arr[ch-'a']++
	}

	for _, ch := range s2 {
		arr[ch-'a']++
	}

	for idx := range arr {
		if arr[idx] == 1 {
			res += string(idx + 'a')
		}
	}
	fmt.Println(res)
}
