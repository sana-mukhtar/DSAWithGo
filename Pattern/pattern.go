package main

import "fmt"

func pattern1(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func pattern2(n int) {
	for i := 1; i < 2*n; i++ {
		var noOfCols int
		if i <= n {
			noOfCols = i
		} else {
			noOfCols = 2*n - i
		}

		for j := 1; j <= noOfCols; j++ {
			fmt.Print("*", " ")
		}
		fmt.Println()
	}
}

func pattern3(n int) {
	for i := 1; i < 2*n; i++ {

		noOfSpaces, noOfStars := 0, 0
		if i <= n {
			noOfSpaces = n - i
			noOfStars = i
		} else {
			noOfSpaces = i - n
			noOfStars = 2*n - i
		}

		for j := 1; j <= noOfSpaces; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= noOfStars; j++ {
			fmt.Print("*", " ")
		}
		fmt.Println()
	}
}

//         1
//       2 1 2
//     3 2 1 2 3
//   4 3 2 1 2 3 4
// 5 4 3 2 1 2 3 4 5
func pattern4(n int) {
	for i := 1; i <= 2*n-1; i++ {
		var noOfCols int
		if i <= n {
			noOfCols = i
		} else {
			noOfCols = 2*n - i
		}
		for j := 1; j <= n-noOfCols; j++ {
			fmt.Print(" ")
		}

		for j := noOfCols; j >= 1; j-- {
			fmt.Print(j, " ")
		}

		for j := 2; j <= noOfCols; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func main() {
	// fmt.Println("Hello World")
	// pattern1(5)
	// pattern2(5)
	// pattern3(5)
	pattern4(5)

}
