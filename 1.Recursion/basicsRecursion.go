package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Print numbers from 1 to n without the help of loops. You only need to complete the function printNos() that takes n as a parameter and prints the number from 1 to n recursively.
func printNos(i, num int) {
	if i > num {
		return
	}
	fmt.Println(i)
	i++
	printNos(i, num)
}

// Print numbers from N to 1 (space separated) without the help of loops.
func printReverseNos(N int) {
	if N < 1 {
		return
	}
	fmt.Println(N)
	N--
	printReverseNos(N)

}

// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// "Given a string, check if the string is palindrome or not."  A string is said to be palindrome if the reverse of the string is the same as the string.
func isPalindrome(str string) bool {
	var filteredStr string
	for _, s := range str {
		if unicode.IsLetter(s) || unicode.IsDigit(s) {
			filteredStr = filteredStr + string(s)
		}
	}
	filteredStr = strings.ToLower(filteredStr)
	endIdx := len(strings.Split(filteredStr, "")) - 1
	return isPalindromeUsingRecursion(filteredStr, 0, endIdx)
}

func isPalindromeUsingRecursion(str string, start, end int) bool {
	if start >= end {
		return true
	}
	if str[start] != str[end] {
		return false
	}
	return isPalindromeUsingRecursion(str, start+1, end-1)
}

// Given an integer n, calculate the sum of series 1^3 + 2^3 + 3^3 + 4^3 + … till n-th term.
// Example:
// Input: n = 5
// Output: 225
// Explanation: 1^3 + 2^3 + 3^3 + 4^3 + 5^3 = 225
func sumOfSeries(n int, sum int) int {
	if n == 0 {
		return sum
	}
	sum += n * n * n
	return sumOfSeries(n-1, sum)
}

func main() {
	// printNos(1, 10)
	// printReverseNos(5)
	// fmt.Println(isPalindrome("0P"))
	// fmt.Println(sumOfSeries(5, 0))
	// fmt.Println(isPalindrome("ABCba"))
	// printN(3)
	// printName(1, 2)
	// sumOfNos(1, 6, 0)
	// fmt.Println(factorialUsingLoop(5))
	// fmt.Println(factorialOfNo(5))
	// fmt.Println(reverseArray([]int{1, 2, 3, 4, 5}))
	// fmt.Println(reverseArrUsingRecursion([]int{1, 2, 3, 4, 5}, 0, 4))
	fmt.Println(checkPalindrome("abaca", 0, 4))
}

// Revision
// print nos till n
var count = 0

func printN(n int) {
	if count > n {
		return
	}
	fmt.Println(count)
	count++
	printN(n)
}

// print name N times
func printName(i, n int) {
	if i > n {
		return
	}
	fmt.Println("Sana")
	printName(i+1, n)
}

func sumOfNos(start, limit, sum int) {
	if start > limit {
		fmt.Println(sum)
		return
	}
	sum = sum + start
	sumOfNos(start+1, limit, sum)
}

// TODO
func factorialOfNo(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorialOfNo(num-1)
}

// i = 0, fact=1*5=5
// i = 1, fact=5*4 =20
// i = 2, fact=20*3=60
// i = 3, fact=60*2=120
// i = 4, fact=120*1=120
func factorialUsingLoop(n int) int {
	fact := 1
	for i := 0; i < n; i++ {
		fact = fact * (n - i) //5//4//3//2
	}
	return fact
}

func reverseArray(originalArr []int) []int {
	reversedArr := make([]int, 0)
	for i := len(originalArr) - 1; i >= 0; i-- {
		reversedArr = append(reversedArr, originalArr[i])
	}
	return reversedArr
}

// [1,2,3,4,5]
func reverseArrUsingRecursion(arr []int, startIdx, endIdx int) []int {
	//base case
	if startIdx > endIdx {
		return arr
	}

	// swap
	temp := arr[startIdx]
	arr[startIdx] = arr[endIdx]
	arr[endIdx] = temp

	return reverseArrUsingRecursion(arr, startIdx+1, endIdx-1)
}

func checkPalindrome(word string, startIdx, endIdx int) bool {
	// base case
	if startIdx > endIdx {
		return true
	}

	w := strings.Split(word, "")
	if w[startIdx] != w[endIdx] {
		return false
	}

	return checkPalindrome(word, startIdx+1, endIdx-1)

}
