package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
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

// [0,0,0,0,1,1,1,2,2,2]
// zeroCount = 4
// oneCount = 3
// twoCount = 3
func sortZeroOnesAndTwos(nums []int) []int {
	zeroCount, oneCount, twoCount := 0, 0, 0
	for _, n := range nums {
		if n == 0 {
			zeroCount++
		} else if n == 1 {
			oneCount++
		} else {
			twoCount++
		}
	}

	for i := 0; i < zeroCount; i++ {
		nums[i] = 0
	}

	for i := zeroCount; i < zeroCount+oneCount; i++ {
		nums[i] = 1
	}

	for i := zeroCount + oneCount; i < len(nums); i++ {
		nums[i] = 2
	}

	return nums
}

func main() {
	// fmt.Println(twoSum([]int{1, 2, 3, 1, 4, 5, 7}, 6))
	// fmt.Println(twoSumUsingTwoPointers([]int{1, 2, 3, 1, 4, 5, 7}, 6))
	fmt.Println(sortZeroOnesAndTwos([]int{1, 2, 0, 0, 0, 2, 2, 2, 1, 1, 1, 2, 2, 0}))
}
