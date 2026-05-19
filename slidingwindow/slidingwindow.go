package main

import "fmt"

func maxSubArraySum(arr []int, k int) int {
	n := len(arr)

	if n < k {
		return -1
	}

	windowSum := 0
	// Reach till the first fixed window of size k
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum
	for i := k; i < n; i++ {
		// slide the window
		windowSum = windowSum - arr[i-k] + arr[i]

		// Do calculation on current window
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	fmt.Println("Start small. Ship something.", maxSubArraySum([]int{2, 1, 5, 1, 3, 2}, 3))
}
