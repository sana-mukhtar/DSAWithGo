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

func firstNegativeOnEachWindow(arr []int, windowSize int) []int {
	n := len(arr)

	if n < windowSize {
		return []int{}
	}

	result, negatives := []int{}, []int{}

	// create first window
	for i := 0; i < windowSize; i++ {
		if arr[i] < 0 {
			negatives = append(negatives, arr[i])
		}
	}

	// do calculation for first window
	if len(negatives) > 0 {
		result = append(result, negatives[0])
	} else {
		result = append(result, 0)
	}

	for i := windowSize; i < len(arr); i++ {
		// Remove first negative if it goes out of window
		if len(negatives) > 0 && arr[i-windowSize] == negatives[0] {
			negatives = negatives[1:]
		}

		// Add incoming negative element
		if arr[i] < 0 {
			negatives = append(negatives, arr[i])
		}

		// Do calculation for current window
		if len(negatives) > 0 {
			result = append(result, negatives[0])
		} else {
			result = append(result, 0)
		}
	}

	return result
}

func countAnagramOccurences(str, ptrn string) int {
	n, k := len(str), len(ptrn)

	if n < k {
		return 0
	}

	count := 0
	patternFreq, windowFreq := make(map[byte]int), make(map[byte]int)

	// Store pattern frequency
	for i := 0; i < k; i++ {
		patternFreq[ptrn[i]]++
	}

	// calculate first window
	for i := 0; i < k; i++ {
		windowFreq[str[i]]++
	}

	// check if first window is anagram
	if isAnagram(patternFreq, windowFreq) {
		count++
	}

	for i := k; i < n; i++ {
		// Remove left character
		windowFreq[str[i-k]]--

		// Add new right character
		windowFreq[str[i]]++

		// 2. Do calculation on current window
		if isAnagram(patternFreq, windowFreq) {
			count++
		}
	}

	return count
}

func isAnagram(a, b map[byte]int) bool {
	for ch, freq := range a {
		if b[ch] != freq {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Start small. Ship something.", maxSubArraySum([]int{2, 1, 5, 1, 3, 2}, 3))
	fmt.Println("negatives in each window", firstNegativeOnEachWindow([]int{-1, -2, 1, 2, 3, -5, 6, -8}, 3))
	fmt.Println("anagram occurences", countAnagramOccurences("forxxorfxdofr", "for"))
	fmt.Println("anagram occurences", countAnagramOccurences("aabaababaa", "abaa"))
}
