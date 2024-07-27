package stringDS

import "fmt"

func CheckForSubsequence() {
	str1 := "gksrek"
	str2 := "geeksforgeeks"
	m, n := len(str1), len(str2)
	i, j := 0, 0

	for i < m && j < n {
		if str1[i] == str2[j] {
			i += 1
		}
		j += 1
	}

	fmt.Println(i == m)
}
