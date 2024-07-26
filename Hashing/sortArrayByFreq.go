package hashing

import (
	"fmt"
	"sort"
)

type Pair struct {
	Value int
	Freq  int
}

func SortByFrequency() {
	arr := []int{5, 5, 4, 6, 4}
	freqMap := make(map[int]int)

	for _, val := range arr {
		freqMap[val]++
	}

	pairs := make([]Pair, 0, len(freqMap))
	for value, freq := range freqMap {
		pairs = append(pairs, Pair{Value: value, Freq: freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Freq == pairs[j].Freq {
			return pairs[i].Value < pairs[j].Value
		}
		return pairs[i].Freq > pairs[j].Freq
	})

	for _, pair := range pairs {
		for j := 0; j < pair.Freq; j++ {
			fmt.Print(pair.Value, " ")
		}
	}
	fmt.Println()
}
