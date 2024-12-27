package main

import (
	"fmt"
	"math"
)

func countElements(arr []int) map[int]int {
	newMap := make(map[int]int)
	for _, val := range arr {
		newMap[val]++
	}
	return newMap
}

// Find the highest/lowest frequency element
func findLowestHighestFrequency(arr []int) (int, int) {
	lowest, highest := math.MaxInt, math.MinInt
	hashmap := countElements(arr)
	lk, hk := 0, 0
	for key, val := range hashmap {
		if val > highest {
			highest = val
			hk = key
		}
		if val < lowest {
			lowest = val
			lk = key
		}
	}
	return lk, hk
}

func main() {
	arr := []int{2, 3, 3, 3, -1, 3, 2, 9, 9}
	// fmt.Println(countElements(arr))
	lowest, highest := findLowestHighestFrequency(arr)
	fmt.Printf("lowest frequency number: %d highest frequency number: %d", lowest, highest)

}
