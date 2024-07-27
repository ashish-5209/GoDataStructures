package stringDS

import "fmt"

func CheckRotatedString() {
	s1 := "geeksforgeeks"
	s2 := "forgeeksgeeks"
	n := len(s1)
	m := len(s2)
	if m > n {
		fmt.Println(false)
	}

	s1 += s1
	s := len(s1)

	for i := 0; i <= s-m; i++ {
		if s1[i:i+m] == s2 {
			fmt.Println(true)
			return
		}
	}
	fmt.Println(false)
}
