package matrix

import "fmt"

func App() {
	matrix := make(map[int]string)
	matrix[0] = "Create Matrix"
	for key, val := range matrix {
		fmt.Println(key, "			", val)
	}
	var val int
	fmt.Scanln(&val)

	switch val {
	case val:
		fmt.Println(matrix[val])
		CreateMatrix()
	}
}
