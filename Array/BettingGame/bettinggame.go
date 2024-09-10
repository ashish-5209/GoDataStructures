package main

import "fmt"

func main() {
	ob := "WLWLLLWWLW"
	bet := 1
	s := 4

	for _, val := range ob {
		if s < bet {
			s = -1
			break
		}
		if string(val) == "W" {
			s += bet
			bet = 1
		} else {
			s -= bet
			bet *= 2
		}
	}

	fmt.Println(s)
}
