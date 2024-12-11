package main

import (
	"fmt"
	"math"
)

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

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0. Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
func reverseInteger(n int) int {
	minNum := math.MinInt32
	maxNum := math.MaxInt32

	reverse := 0
	for n != 0 {
		lastDigit := n % 10
		n = n / 10
		if reverse > maxNum/10 || (reverse == maxNum/10 && lastDigit > 7) {
			return 0
		}
		if reverse < minNum/10 || (reverse == minNum/10 && lastDigit < -8) {
			return 0
		}
		reverse = reverse*10 + lastDigit
	}

	return reverse
}

func main() {
	fmt.Println("hello")
	fmt.Println(countDigits(3336999))
	fmt.Print(reverseInteger(123))
}
