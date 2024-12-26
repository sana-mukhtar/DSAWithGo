package main

import "fmt"

func countElements(arr []int) map[int]int {
	newMap := make(map[int]int)
	for _, val := range arr {
		newMap[val]++
	}
	return newMap
}

func main() {
	arr := []int{2, 3, 5, 2, 4, 5, 1, 2, 9}
	fmt.Println(countElements(arr))

}
