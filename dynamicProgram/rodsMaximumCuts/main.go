package main

import "fmt"

type Element struct {
	n, a, b, c int
}

type ElementI interface {
	GetMaximumCutsRecursion() int
	GetMaximumCutsOptimized() int
}

func newElem(n, a, b, c int) ElementI {
	return Element{
		n: n,
		a: a,
		b: b,
		c: c,
	}
}

func (e Element) GetMaximumCutsRecursion() int {
	if e.n < 0 {
		return -1
	}
	if e.n == 0 {
		return 0
	}
	e.n -= e.a
	val1 := e.GetMaximumCutsRecursion()

	e.n += e.a
	e.n -= e.b
	val2 := e.GetMaximumCutsRecursion()

	e.n += e.b
	e.n -= e.c
	val3 := e.GetMaximumCutsRecursion()

	res := max(val1, val2, val3)

	if res == -1 {
		return res
	}
	return res + 1
}

func (e Element) GetMaximumCutsOptimized() int {
	dp := make([]int, (e.n + 1))
	for i := 1; i < (e.n + 1); i++ {
		dp[i] = -1
		if (i - e.a) >= 0 {
			dp[i] = max(dp[i], dp[i-e.a])
		}
		if (i - e.b) >= 0 {
			dp[i] = max(dp[i], dp[i-e.b])
		}
		if (i - e.c) >= 0 {
			dp[i] = max(dp[i], dp[i-e.c])
		}

		if dp[i] != -1 {
			dp[i]++
		}
	}
	return dp[e.n]
}

func main() {
	obj := newElem(17, 10, 11, 3)
	fmt.Println(obj.GetMaximumCutsRecursion())
	fmt.Println(obj.GetMaximumCutsOptimized())
}
