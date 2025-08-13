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

// i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
func threeSum(nums []int) [][]int {
	final := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := i + 2; k < len(nums); k++ {
				if (i != j || i != k || j != k) && (nums[i]+nums[j]+nums[k] == 0) {
				}
			}
		}
	}

	return final
}

func pascalsTriangle(numRows int) [][]int {
	result := [][]int{}

	for i := 0; i < numRows; i++ {
		// Create a row with i+1 elements
		row := make([]int, i+1)

		// First and last elements are always 1
		row[0], row[i] = 1, 1

		// Fill the middle elements of the row with sum of above-left + above-right from the previous row
		for j := 1; j < i; j++ {
			aboveLeft := result[i-1][j-1]  // element above-left
			aboveRight := result[i-1][j]   // element directly above

			row[j] = aboveLeft + aboveRight
		}

		// Add the completed row to the result
		result = append(result, row)
	}

	return result
}


func main() {
	fmt.Println(majorityElement2([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
