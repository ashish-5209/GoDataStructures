package matrix

import "fmt"

func CreateMatrix() {
	var rows, cols int

	fmt.Print("Enter the no of rows: ")
	fmt.Scanln(&rows)
	fmt.Print("Enter the no of columns: ")
	fmt.Scanln(&cols)

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	fmt.Println("Enter the matrix values row by rows:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter the value of matrix[%d][%d]: ", i, j)
			fmt.Scanln(&matrix[i][j])
		}
	}

	fmt.Println("The matrix is: ")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
