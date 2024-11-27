package main

import "fmt"

type Element struct {
	string1, string2 string
}

type ElementI interface {
	LCSRecursion(m, n int) int
	LCSDP(m, n int) int
}

func newElem(string1, string2 string) ElementI {
	return &Element{
		string1: string1,
		string2: string2,
	}
}

func (e Element) LCSRecursion(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if e.string1[m-1] == e.string2[n-1] {
		return 1 + e.LCSRecursion(m-1, n-1)
	}

	return max(e.LCSRecursion(m, n-1), e.LCSRecursion(m-1, n))
}

func (e Element) LCSDP(m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if e.string1[i-1] == e.string2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	string1 := "XYZW"
	string2 := "XYWZ"

	obj := newElem(string1, string2)
	fmt.Println(obj.LCSRecursion(len(string1), len(string2)))
	fmt.Println(obj.LCSDP(len(string1), len(string2)))
}
