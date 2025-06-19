package main

import "fmt"

func majorityElement2(nums []int) []int {
	result := []int{}

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}

		if count > len(nums)/3 {
			result = append(result, nums[i])
		}
	}

	return result
}

func main() {
	fmt.Println(majorityElement2([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
