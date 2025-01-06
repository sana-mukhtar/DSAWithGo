package main

import (
	"fmt"
	"math"
)

func findLargest(arr []int) int {
	max := math.MinInt
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func findSecondLargest(arr []int) int {
	largest := math.MinInt
	secondLargest := math.MinInt
	for _, val := range arr {
		if val > largest {
			secondLargest = largest
			largest = val
		}
	}
	return secondLargest
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	max := findLargest([]int{4, 52, 2, 78, 55})
	fmt.Println(max)
	fmt.Println(findSecondLargest([]int{3, 45, 32, 78, 12, 99, 100}))
	fmt.Println(isSorted([]int{2, 3, 4, 5, 5, 6}))
}
