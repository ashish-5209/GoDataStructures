package stringDS

import (
	"fmt"
	"strconv"
	"unicode"
)

func SumOfNumberInString() {
	s := "1abc23"
	res := 0
	currentNum := ""

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			currentNum += string(ch)
		} else {
			if currentNum != "" {
				num, err := strconv.Atoi(currentNum)
				if err == nil {
					res += num
				}
				currentNum = ""
			}
		}
	}

	if currentNum != "" {
		num, err := strconv.Atoi(currentNum)
		if err == nil {
			res += num
		}
		currentNum = ""
	}

	fmt.Println(res)
}
