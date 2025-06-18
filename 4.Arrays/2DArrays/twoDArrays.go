package main

import "fmt"

func rowWiseTraversal(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print("row:", i, " ")
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func columnWiseTraversal(matrix [][]int) {
	for j := 0; j < len(matrix[0]); j++ {
		fmt.Print("col", j, " ")
		for i := 0; i < len(matrix); i++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func reverseRowWiseTraversal(matrix [][]int) {
	for i := len(matrix) - 1; i >= 0; i-- {
		fmt.Print("row:", i, " ")
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func reverseColumnWiseTraversal(matrix [][]int) {
	for j := len(matrix) - 1; j >= 0; j-- {
		fmt.Print("reverse column wise", j, " ")
		for i := 0; i < len(matrix); i++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

// Transposing a matrix means flipping it over its diagonal — rows become columns and columns become rows.
func transposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	columns := len(matrix[0])
	result := make([][]int, columns) // because no of columns is going to be length of row for transposed matrix

	for j := 0; j < columns; j++ {
		result[j] = make([]int, rows)
		for i := 0; i < rows; i++ {
			// Transpose: element at (i, j) in original becomes (j, i) in transposed
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

// *
// * *
// * * *
// * * * *
// * * * * *
func pattern() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}

func generatePascal(n int) [][]int {
	result := [][]int{}

	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result = append(result, row)
	}

	return result
}

func main() {
	tri := generatePascal(5)
	printMatrix(tri)
}
