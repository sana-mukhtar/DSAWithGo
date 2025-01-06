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

func main() {
	max := findLargest([]int{4, 52, 2, 78, 55})
	fmt.Println(max)
	fmt.Println(findSecondLargest([]int{3, 45, 32, 78, 12, 99, 100}))
}
