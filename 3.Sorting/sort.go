package main

import "fmt"

// Problem Statement: Given an array of N integers, write a program to implement the Selection sorting algorithm.
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		lowestNum := arr[i]
		track := i
		for j := 0 + i; j < len(arr); j++ {
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

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		for j := i - 1; j >= 0 && key < arr[j]; j-- {
			arr[j+1] = arr[j]
			arr[j] = key
		}
	}
	return arr
}

// 1 loop on array -> compare i and i+1-> if i+1 is smaller -> iterate inversely

func main() {
	// 	fmt.Println("selectionsort:", selectionSort([]int{64, 25, 12, 22, 11}))
	// 	fmt.Println("bubblesort:", bubbleSort([]int{38, 27, 43, 3, 9, 82, 10}))
	// 	fmt.Println("bubblesort:", bubbleSort([]int{64, 25, 12, 22, 11}))
	fmt.Println("insertionsort:", insertionSort([]int{5, 4, 3, 2, 7}))
	// for j := 0; j >= 0; j-- {
	// 	println(j)
	// }

}
