package main

import "fmt"

func printPattern(n int) {
	fmt.Print(n, " ")
	if n <= 0 {
		return
	}
	printPattern(n - 5)
	fmt.Print(n, " ")
}

func main() {
	n := 16
	printPattern(n)
	fmt.Println()
}
