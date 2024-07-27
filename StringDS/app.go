package stringDS

import "fmt"

func App() {
	str := make(map[int]string)
	str[0] = "Count No Of One in Binary String"
	str[1] = "First Occurence of one string in another string"
	str[2] = "Chcek if string in rotated"
	str[3] = "Chcek for isomorphic string"
	str[4] = "Chcek for subsequence"
	str[5] = "Maximum Occuring character"
	str[6] = "Case Specific string sorting"
	str[7] = "Check if string in anagram"

	for key, val := range str {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)
	switch val {
	case 0:
		fmt.Println(str[val])
		CountNoOfOneCombination()
	case 1:
		fmt.Println(str[val])
		GetFirstOccurrenceOfString()
	case 2:
		fmt.Println(str[val])
		CheckRotatedString()
	case 3:
		fmt.Println(str[val])
		CheckIsomorphicString()
	case 4:
		fmt.Println(str[val])
		CheckForSubsequence()
	case 5:
		fmt.Println(str[val])
		MaximumOccurrenceCharacter()
	case 6:
		fmt.Println(str[val])
		CaseSpecificStringSorting()
	case 7:
		fmt.Println(str[val])
		CheckAnagram()
	}
}
