package stringDS

import "fmt"

func ModifyString() {
	s := "aaaaa"
	count := 0
	i := 0

	for i < len(s)-2 {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			count++
			i += 2
		} else if s[i] == s[i+1] && s[i] != s[i+2] {
			i += 2
		} else {
			i += 1
		}
	}

	fmt.Println(count)
}
