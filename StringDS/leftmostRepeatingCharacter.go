package stringDS

import "fmt"

func LeftMostRepeatingCharacter() {
	s := "abcd"
	var arr [26]int

	for i := range arr {
		arr[i] = -1
	}
	res := -1

	for i := len(s) - 1; i >= 0; i-- {
		if arr[s[i]-'a'] == -1 {
			arr[s[i]-'a'] = i
		} else {
			res = i
		}
	}

	fmt.Println(res)
}
