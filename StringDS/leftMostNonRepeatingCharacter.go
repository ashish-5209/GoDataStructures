package stringDS

import (
	"fmt"
	"math"
)

func LeftNonMostRepeatingCharacter() {
	s := "zxvczbtxyzvy"
	var arr [26]int
	for i := range arr {
		arr[i] = -1
	}

	for idx, ch := range s {
		if arr[ch-'a'] == -1 {
			arr[ch-'a'] = idx
		} else {
			arr[ch-'a'] = -2
		}
	}

	res := math.MaxInt32

	for _, val := range arr {
		if val >= 0 {
			res = min(res, val)
		}
	}

	if res == math.MaxInt32 {
		fmt.Println("$")
	} else {
		fmt.Println(string(s[res]))
	}
}
