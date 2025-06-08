package main

import (
	"fmt"
	"sort"
)

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

func sortZeroOnesAndTwosUsingDNF(nums []int) []int {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		if nums[mid] == 0 {
			// swap nums[mid],low
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			// swap nums[mid],high
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}

	return nums
}

// 0,1,2,1,2,1
func majorityElement(nums []int) int {
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}

func majorityElementBetter(nums []int) int {
	sort.Ints(nums)

	freq, ans := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if ans == nums[i-1] {
			freq++
			ans = nums[i]
			if freq > len(nums)/2 {
				return nums[i]
			}
		} else {
			freq = 1
		}
	}

	return 0
}

func main() {
	// fmt.Println(twoSum([]int{1, 2, 3, 1, 4, 5, 7}, 6))
	// fmt.Println(twoSumUsingTwoPointers([]int{1, 2, 3, 1, 4, 5, 7}, 6))
	// fmt.Println(sortZeroOnesAndTwos([]int{1, 2, 0, 0, 0, 2, 2, 2, 1, 1, 1, 2, 2, 0}))
	// fmt.Println(sortZeroOnesAndTwosUsingDNF([]int{1, 2, 0, 0, 0, 2, 2, 2, 1, 1, 1, 2, 2, 0}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElementBetter([]int{2, 2, 1, 1, 1, 2, 2}))
}
