package main

import "fmt"

// 1. Selection Sort – “Find the smallest and put it first”
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

// In Bubble Sort, for every iteration:
// Compare the current element with the next one (arr[i] with arr[i+1]).
// If the current element is greater, swap them.
// After each pass, the largest value "bubbles" to the correct position at the end.
// Keep repeating this process until the whole array is sorted.
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

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	sorted := []int{}
	leftptr, rightptr := 0, 0

	for leftptr < len(left) && rightptr < len(right) {
		if left[leftptr] <= right[rightptr] {
			sorted = append(sorted, left[leftptr])
			leftptr++
		} else {
			sorted = append(sorted, right[rightptr])
			rightptr++
		}
	}
	sorted = append(sorted, left[leftptr:]...)
	sorted = append(sorted, right[rightptr:]...)
	return sorted
}

func main() {
	// 	fmt.Println("selectionsort:", selectionSort([]int{64, 25, 12, 22, 11}))
	// 	fmt.Println("bubblesort:", bubbleSort([]int{38, 27, 43, 3, 9, 82, 10}))
	// 	fmt.Println("bubblesort:", bubbleSort([]int{64, 25, 12, 22, 11}))
	// fmt.Println("insertionsort:", insertionSort([]int{5, 4, 3, 2, 7}))

	// fmt.Println(mergeSort([]int{2, 8, 4, 5, 9, 1, 12, 5}))
	a := selSort([]int{1, 2, 4, 5, 8, 0})
	b := bubbleSort2([]int{2, 3, 1, 4})
	c := insertionSort2([]int{1, 4, 7, 3})
	fmt.Println(a, b, c)

}

// step 1 - find min element
func selSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}

		if minIdx != i {
			temp := arr[i]
			arr[i] = arr[minIdx]
			arr[minIdx] = temp
		}
	}
	return arr
}

func bubbleSort2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				//swap
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

// TODO
func insertionSort2(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		prev := i - 1
		if prev > curr && prev >= 0 {
			// start shifting
			arr[prev+1] = arr[prev]
			prev--
		}
		arr[prev+1] = curr
	}
	return arr
}
