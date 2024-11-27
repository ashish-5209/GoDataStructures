package main

import "fmt"

type Element struct {
	str1, str2 string
}

type ElementI interface {
	GetMinimumWindow() string
}

func newElem(str1, str2 string) ElementI {
	return Element{
		str1: str1,
		str2: str2,
	}
}

func (e Element) GetMinimumWindow() string {
	str2Len := len(e.str2)
	str2Map := make(map[string]int)
	for _, val := range e.str2 {
		str2Map[string(val)]++
	}
	res := ""
	start, end, count := 0, 0, len(e.str1)+1
	for end < len(e.str1) {
		str2Map[string(e.str1[end])]--
		if str2Map[string(e.str1[end])] >= 0 {
			str2Len--
		}

		for str2Len == 0 && start <= end {
			if count > end-start+1 {
				count = end - start + 1
				res = e.str1[start : end+1]
			}
			str2Map[string(e.str1[start])]++
			if str2Map[string(e.str1[start])] > 0 {
				str2Len++
			}
			start++
		}
		end++
	}
	return res
}

func main() {
	obj := newElem("ADOBECODEBANC", "ABC")
	fmt.Println((obj.GetMinimumWindow()))
}
