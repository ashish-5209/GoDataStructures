package stringDS

import (
	"fmt"
	"strings"
)

func ReverseWords() {
	s := "i.like.this.program.very.much"
	words := strings.Split(s, ".")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(strings.Join(words, "."))
}
