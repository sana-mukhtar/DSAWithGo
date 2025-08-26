package main

import "fmt"

// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
func majorityElementII(nums []int) []int {
    // This will store all elements that appear more than n/3 times
    result := []int{}

    // Outer loop: pick each element one by one
    for i := 0; i < len(nums); i++ {
        count := 0 // to count occurrences of nums[i]

        // Skip if nums[i] is already in the result (to avoid duplicates)
        alreadyExists := false
        for _, val := range result {
            if val == nums[i] {
                alreadyExists = true
                break
            }
        }
        if alreadyExists {
            continue
        }

        // Inner loop: count how many times nums[i] appears from i to end
        for j := i; j < len(nums); j++ {
            if nums[i] == nums[j] {
                count++
            }
        }

        // If nums[i] occurs more than n/3 times, add it to the result
        if count > len(nums)/3 {
            result = append(result, nums[i])
        }
    }

    // Return all elements that appeared more than n/3 times
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

// Row 0:        1
// Row 1:       1 1
// Row 2:      1 2 1
// Row 3:     1 3 3 1
// Row 4:    1 4 6 4 1
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
	fmt.Println(majorityElementII([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
