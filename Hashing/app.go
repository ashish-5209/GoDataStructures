package hashing

import "fmt"

func App() {
	hashing := make(map[int]string)
	hashing[0] = "Pair Sum Distinct Elements"
	hashing[1] = "Count Non Repeating Elements"
	hashing[2] = "Sort array by frequency"
	hashing[3] = "Sort array with respect to other array"
	hashing[4] = "First repeating element"
	hashing[5] = "Union of arrays"
	hashing[6] = "Positive Negative pair"
	hashing[7] = "SubArray with zero sum"
	hashing[8] = "SubArray with given sum"
	hashing[9] = "Longest SubArray with given sum"
	hashing[10] = "Longest consecutive subsequence"
	hashing[11] = "Account Merge"

	for key, val := range hashing {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)
	switch val {
	case 0:
		fmt.Println(hashing[val])
		PairSumDistinct()
	case 1:
		fmt.Println(hashing[val])
		CountNonRepeatingElem()
	case 2:
		fmt.Println(hashing[val])
		SortByFrequency()
	case 3:
		fmt.Println(hashing[val])
		SortArrayByOtherArrray()
	case 4:
		fmt.Println(hashing[val])
		FirstRepeatingElement()
	case 5:
		fmt.Println(hashing[val])
		UnionOfArray()
	case 6:
		fmt.Println(hashing[val])
		FindPositiveNegativePair()
	case 7:
		fmt.Println(hashing[val])
		SubArrayWithZeroSum()
	case 8:
		fmt.Println(hashing[val])
		SubArrayWithGivenSum()
	case 9:
		fmt.Println(hashing[val])
		LongestSubarraySum()
	case 10:
		fmt.Println(hashing[val])
		LongestConsecutiveSubsequence()
	case 11:
		fmt.Println(hashing[val])
		AccountMerge()
	}
}
