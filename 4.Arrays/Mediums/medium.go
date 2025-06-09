package main

import (
	"fmt"
	"sort"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Approach:
// Loop through each pair of indices (i, j)
// Check if nums[i] + nums[j] == k.
// If yes, return [i, j].
// If no pair found, return [].
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
	sort.Ints(nums) // tc - O(nlogn)

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

// using Moore's voting algorithm
func majorityElementOptimal(nums []int) int {
	freq, ans := 0, 0

	for i := 0; i < len(nums); i++ {
		if freq == 0 {
			ans = nums[i]
		}
		if nums[i] == ans {
			freq++
		} else {
			freq--
		}
	}

	return ans
}

// brute force, leetcode's some test failing using this TC O(n2)
func maxSubArray(nums []int) int {
	maxSum := 0
	for start := 0; start < len(nums); start++ {
		currSum := 0
		for end := start; end < len(nums); end++ {
			currSum += nums[end]
			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}

	return maxSum
}

// optimal, TC O(n)
func maxSubArrayUsingKadanesAlgo(nums []int) int {
	currSum, maxSum := 0, nums[0]
	for start := 0; start < len(nums); start++ {
		currSum += nums[start]
		if currSum > maxSum {
			maxSum = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum
}

// optimal, TC O(n)
func printMaxSubArrayUsingKadanesAlgo(nums []int) {
	currSum, maxSum, start, tempStart, end := 0, nums[0], 0, 0, 0
	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		if currSum > maxSum {
			maxSum = currSum
			start = tempStart
			end = i
		}
		if currSum < 0 {
			currSum = 0
			tempStart = i + 1
		}
	}

	fmt.Println("MaxSubArray", nums[start:end+1])
}

func main() {
	printMaxSubArrayUsingKadanesAlgo([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}
