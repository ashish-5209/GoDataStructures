package stringDS

import "fmt"

func GetFirstOccurrenceOfString() {
	str1 := "GeeksForGeeks"
	str2 := "For"

	n := len(str1)
	m := len(str2)

	if m == 0 || m > n {
		fmt.Println(-1)
		return
	}

	for i := 0; i <= n-m; i++ {
		if str1[i:i+m] == str2 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
