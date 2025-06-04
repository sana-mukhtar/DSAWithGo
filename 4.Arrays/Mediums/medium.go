package main

import "fmt"

func twoSum(nums []int, k int) []int {
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == k {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSumUsingTwoPointers(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start < end {
		sum := nums[start] + nums[end]
		
		if sum == target {
			return []int{start, end}
		} else if sum < target {
			start++
		} else {
			end--
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 1, 4, 5, 7}, 6))
	fmt.Println(twoSumUsingTwoPointers([]int{1, 2, 3, 1, 4, 5, 7}, 6))
}
