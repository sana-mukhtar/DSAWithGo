package main

import "fmt"

// Q.1 Given a positive integer n, count the number of digits in n that divide n evenly (i.e., without leaving a remainder). Return the total number of such digits.
func countDigits(n int) int {
	dividors := 0
	copy := n
	for n > 0 {
		ld := n % 10
		if copy%ld == 0 {
			dividors++
		}
		n = n / 10
	}
	return dividors
}

func main() {
	fmt.Println("hello")
	fmt.Println(countDigits(3336999))
}
