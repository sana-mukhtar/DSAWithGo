package main

import "fmt"

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

//The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// printNos(1, 10)
	printReverseNos(5)
}
