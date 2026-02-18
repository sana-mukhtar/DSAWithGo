package main

import "fmt"

// 1. Selection Sort – “Find the smallest and put it first”
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		minIdx := i
		for j := 0 + i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIdx = j // change minIdx to the lowest num
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
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
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// Insertion sort places each element into its correct position in the already sorted part of the array
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Shift all elements in the sorted part (0..i-1)
		// that are greater than 'key' one position to the right
		// to create the correct position for 'key'
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Insert 'key' at its correct position
		arr[j+1] = key
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])

	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	sortedArr := []int{}
	leftIdx, rightIdx := 0, 0

	for leftIdx < len(leftArr) && rightIdx < len(rightArr) {
		if leftArr[leftIdx] < rightArr[rightIdx] {
			sortedArr = append(sortedArr, leftArr[leftIdx])
			leftIdx++
		} else {
			sortedArr = append(sortedArr, rightArr[rightIdx])
			rightIdx++
		}
	}

	sortedArr = append(sortedArr, leftArr[leftIdx:]...)
	sortedArr = append(sortedArr, rightArr[rightIdx:]...)

	return sortedArr
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// get pivot index
		p := partition(arr, low, high)

		// sort left part
		quickSort(arr, low, p-1)

		// sort right part
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // choose last element as pivot
	i := low - 1

	// whenever element smaller than pivot is found swap with i
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// place pivot at correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	// fmt.Println("selectionsort:", selectionSort([]int{64, 25, 12, 22, 11}))
	// fmt.Println("bubblesort:", bubbleSort([]int{38, 27, 43, 3, 9, 82, 10}))
	// 	fmt.Println("bubblesort:", bubbleSort([]int{64, 25, 12, 22, 11}))
	// fmt.Println("insertionsort:", insertionSort([]int{5, 4, 3, 2, 7}))

	fmt.Println(mergeSort([]int{2, 8, 4, 5, 9, 1, 12, 5}))
	arr := []int{64, 25, 12, 22, 11, 2, 22}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("quicksort:", arr)

	// a := selSort([]int{1, 2, 4, 5, 8, 0})
	// b := bubbleSort2([]int{2, 3, 1, 4})
	// c := insertionSort2([]int{1, 4, 7, 3})
	// fmt.Println(a, b, c)

}
