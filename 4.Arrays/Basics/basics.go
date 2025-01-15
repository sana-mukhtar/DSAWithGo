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

func isSortedAsc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func isSortedDesc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

func isSorted(arr []int) bool {
	if arr[0] > arr[1] {
		return isSortedDesc(arr)
	} else {
		return isSortedAsc(arr)
	}
}

func removeDuplicates(nums []int) int {
	var uniqueArray []int
	uniqueArray = append(uniqueArray, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		uniqueArray = append(uniqueArray, nums[i])
	}
	fmt.Println("unique array: ", uniqueArray)
	return len(uniqueArray)
}

func rotate(nums []int, k int) {
	length := len(nums) - 1  // 6
	for i := 0; i < k; i++ { // 3 steps
		temp := nums[length]
		for j := length; j > 0; j-- {
			fmt.Println(nums[j], ":", nums[j-1])
			nums[j] = nums[j-1]
		}
		nums[0] = temp
		fmt.Println(i, ":", nums)
	}
}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			temp := nums[i]
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[len(nums)-1] = temp
			fmt.Println(i, ":",nums)
		} else {
			continue
		}
	}
}

func main() {
	// max := findLargest([]int{4, 52, 2, 78, 55})
	// fmt.Println(max)
	// fmt.Println(findSecondLargest([]int{3, 45, 32, 78, 12, 99, 100}))
	// fmt.Println(isSorted([]int{2, 3, 4, 5, 5, 6}))
	// fmt.Println(isSorted([]int{20, 18, 17, 16, 10, 7, 0}))
	// fmt.Println(isSorted([]int{2, 3, 4, 5, 15, 6}))
	// fmt.Println(removeDuplicates([]int{1,1,3, 3, 4, 5, 7, 7, 8}))
	// rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	moveZeroes([]int{0,1,0,3,12})
}
