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

// Q.2 Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0. Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
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

// Given an integer x, return true if x is a palindrome, and false otherwise. Example 1:Input: x = 121 Output: trueExplanation: 121 reads as 121 from left to right and from right to left. Example 2:Input: x = -121 Output: falseExplanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	rev := 0
	lastDigit := 0
	copyNum := x
	for x != 0 {
		lastDigit = x % 10
		rev = rev*10 + lastDigit
		x = x / 10
	}
	if rev == copyNum {
		return true
	}
	return false
}

// Q.4 Given two integers a and b, write a function lcmAndGcd() to compute their LCM and GCD. The function inputs two integers a and b and returns a list containing their LCM and GCD.
func lcmAndGcd(n1, n2 int) (int, int) {
	copyn1 := n1
	copyn2 := n2
	for n1%n2 != 0 {
		rem := n1 % n2
		n1 = n2
		n2 = rem

	}
	gcd := n2
	lcm := (copyn1 * copyn2) / gcd
	return lcm, gcd
}

// Q.5 Given an integer n, write a function to determine if it is an Armstrong number. An Armstrong number (also known as a Narcissistic number) for a number n is a number that is equal to the sum of its own digits each raised to the power of the number of digits.
func isArmstrongNumber(n int) bool {
	return false
}

func main() {
	fmt.Println("hello")
	fmt.Println(countDigits(3336999))
	fmt.Println(reverseInteger(123))
	fmt.Println(lcmAndGcd(14, 8))
}
