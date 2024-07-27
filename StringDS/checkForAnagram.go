package stringDS

import "fmt"

func CheckAnagram() {
	s1 := "geeksforgeeks"
	s2 := "forgeeksgeeks"
	if len(s1) != len(s2) {
		fmt.Println(false)
		return
	}
	// arr := make([]int, 26)
	res := 0

	for idx, ch := range s1 {
		char1 := rune(s2[idx])
		res ^= int(ch) ^ int(char1)
	}

	fmt.Println(res == 0)
}
