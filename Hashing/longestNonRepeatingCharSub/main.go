package main

import "fmt"

type Element struct {
	str string
}

type ElementI interface {
	GetLongestNonReapeatingCharacterSubstring() int
}

func newElem(str string) ElementI {
	return Element{
		str: str,
	}
}

func (e Element) GetLongestNonReapeatingCharacterSubstring() int {
	count, start, end := 0, 0, 0
	var charString [255]int
	for end < len(e.str) {
		for charString[e.str[end]] > 0 {
			charString[e.str[start]]--
			start++
		}
		charString[e.str[end]]++
		end++
		count = max(count, end-start)
	}
	return count
}

func main() {
	obj := newElem("pwwkew")
	fmt.Println(obj.GetLongestNonReapeatingCharacterSubstring())
}
