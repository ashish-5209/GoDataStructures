package main

import "fmt"

type Element struct {
	string1, string2 string
}

type ElementI interface {
	EditDistanceRecursion(m, n int) int
	EditDistanceDP(m, n int) int
}

func newElem(string1, string2 string) ElementI {
	return &Element{
		string1: string1,
		string2: string2,
	}
}

func (e Element) EditDistanceRecursion(m, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if e.string1[m-1] == e.string2[n-1] {
		return e.EditDistanceRecursion(m-1, n-1)
	}

	return 1 + min(e.EditDistanceRecursion(m-1, n-1), e.EditDistanceRecursion(m-1, n), e.EditDistanceRecursion(m, n-1))
}

func (e Element) EditDistanceDP(m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m + 1 {
		for j := range n + 1 {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}

			if e.string1[i-1] == e.string2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func main() {
	string1, string2 := "GEEXSFRGEEKKS", "GEEKSFORGEEKS"
	obj := newElem(string1, string2)
	fmt.Println(obj.EditDistanceRecursion(len(string1), len(string2)))
	fmt.Println(obj.EditDistanceDP(len(string1), len(string2)))
}
