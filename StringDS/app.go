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
	str[8] = "Left most repeating character"
	str[9] = "Left most non repeating character"
	str[10] = "Modify string"
	str[11] = "Reverse string"
	str[12] = "Remove and Concanate String"
	str[13] = "Sum of digit in String"
	str[14] = "Find Nth Prime Digit Number"
	str[15] = "Telehone Problem"

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
	case 8:
		fmt.Println(str[val])
		LeftMostRepeatingCharacter()
	case 9:
		fmt.Println(str[val])
		LeftNonMostRepeatingCharacter()
	case 10:
		fmt.Println(str[val])
		ModifyString()
	case 11:
		fmt.Println(str[val])
		ReverseWords()
	case 12:
		fmt.Println(str[val])
		RemoveAndConcanateString()
	case 13:
		fmt.Println(str[val])
		SumOfNumberInString()
	case 14:
		fmt.Println(str[val])
		FindNthPrimeDigitNumber()
	case 15:
		fmt.Println(str[val])
		TelephoneProblem()
	}
}
