package main

import (
	"fmt"
)

// Problem Statement: Given an array of N integers, write a program to implement the Selection sorting algorithm.
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		lowestNum := arr[i]
		track := i
		for j := 0+i; j < len(arr); j++ {
			if arr[j] < lowestNum {
				lowestNum = arr[j]
				track = j
			}
		}
		temp := arr[i]
		arr[i] = lowestNum
		arr[track] = temp
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{64, 25, 12, 22, 11}))
	fmt.Println(selectionSort([]int{38, 27, 43, 3, 9, 82, 10}))
}
