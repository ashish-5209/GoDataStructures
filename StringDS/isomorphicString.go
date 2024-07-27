package stringDS

import "fmt"

func CheckIsomorphicString() {
	str1 := "aab"
	str2 := "xxy"

	if len(str1) != len(str2) {
		fmt.Println(false)
		return
	}
	mapStr1 := make(map[rune]rune)
	mapStr2 := make(map[rune]rune)

	for i, char1 := range str1 {
		char2 := rune(str2[i])
		fmt.Println("rune", char1, char2)

		if ch, ok := mapStr1[char1]; ok {
			if ch != char2 {
				fmt.Println(false)
				return
			}
		} else {
			if ch, ok := mapStr2[char2]; ok {
				if ch != char1 {
					fmt.Println(false)
					return
				}
			}
		}

		mapStr1[char1] = char2
		mapStr2[char2] = char1
	}
	fmt.Println(true)
}
