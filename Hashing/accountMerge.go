package hashing

import (
	"fmt"
	"sort"
)

type DSU struct {
	parent, rank []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent, rank}
}

func (dsu *DSU) find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.find(dsu.parent[x])
	}
	return dsu.parent[x]
}

func (dsu *DSU) union(x, y int) {
	rootX := dsu.find(x)
	rootY := dsu.find(y)
	if rootX != rootY {
		if dsu.rank[rootX] > dsu.rank[rootY] {
			dsu.parent[rootY] = rootX
		} else if dsu.rank[rootX] < dsu.rank[rootY] {
			dsu.parent[rootX] = rootY
		} else {
			dsu.parent[rootY] = rootX
			dsu.rank[rootX]++
		}
	}
}

func accountsMerge(accounts [][]string) [][]string {
	emailToIndex := make(map[string]int)
	emailToName := make(map[string]string)

	// DSU to manage the connected components
	dsu := newDSU(len(accounts) * 10)

	// Assign a unique index to each email and store the corresponding owner
	idx := 0
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, found := emailToIndex[email]; !found {
				emailToIndex[email] = idx
				emailToName[email] = name
				idx++
			}
			dsu.union(emailToIndex[account[1]], emailToIndex[email])
		}
	}

	// Group emails by their root index in the DSU
	indexToEmails := make(map[int][]string)
	for email, index := range emailToIndex {
		root := dsu.find(index)
		indexToEmails[root] = append(indexToEmails[root], email)
	}

	// Create the result by sorting and formatting the emails
	var result [][]string
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		name := emailToName[emails[0]]
		mergedAccount := append([]string{name}, emails...)
		result = append(result, mergedAccount)
	}
	return result
}

func AccountMerge() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}

	mergedAccounts := accountsMerge(accounts)
	for _, account := range mergedAccounts {
		fmt.Println(account)
	}
}
