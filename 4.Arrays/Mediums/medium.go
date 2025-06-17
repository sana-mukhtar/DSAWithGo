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

// TC-O(n)
func maxProfitBuyAndSellStock(prices []int) int {
	maxProfit, bestBuy := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if bestBuy > prices[i] {
			bestBuy = prices[i]
		}

		profit := prices[i] - bestBuy
		if maxProfit < profit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.
// You should return the array of nums such that the the array follows the given conditions:
// Every consecutive pair of integers have opposite signs.
// For all integers with the same sign, the order in which they were present in nums is preserved.
// The rearranged array begins with a positive integer.
// Return the modified array after rearranging the elements to satisfy the aforementioned conditions.
// Input: nums = [3,1,-2,-5,2,-4]
// Output: [3,-2,1,-5,2,-4]
func rearrangeArray(nums []int) []int {
	rearrangedArray, posNums, negNums := []int{}, []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			posNums = append(posNums, nums[i])
		}
		if nums[i] < 0 {
			negNums = append(negNums, nums[i])
		}
	}

	for i := 0; i < len(nums)/2; i++ {
		rearrangedArray = append(rearrangedArray, posNums[i])
		rearrangedArray = append(rearrangedArray, negNums[i])
	}

	return rearrangedArray
}

// Leetcode 88. You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n
func mergeBruteForce(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	mergedNums := make([]int, 0, m+n)
	firstArrIdx, secondArrIdx := 0, 0

	for firstArrIdx < m && secondArrIdx < n {
		if nums1[firstArrIdx] < nums2[secondArrIdx] {
			mergedNums = append(mergedNums, nums1[firstArrIdx])
			firstArrIdx++
		} else {
			mergedNums = append(mergedNums, nums2[secondArrIdx])
			secondArrIdx++
		}
	}

	// Append remaining elements
	for firstArrIdx < m {
		mergedNums = append(mergedNums, nums1[firstArrIdx])
		firstArrIdx++
	}

	for secondArrIdx < n {
		mergedNums = append(mergedNums, nums2[secondArrIdx])
		secondArrIdx++
	}

	// Copy back into nums1
	for i := 0; i < len(mergedNums); i++ {
		nums1[i] = mergedNums[i]
	}
}

// TODO: fix this
func countSubarrayWithSumK(nums []int, k int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		sum = nums[i] + nums[i+1]
		if sum == k || nums[i] == k {
			count++
		}
	}

	return count
}

func longestConsecutiveBruteForce(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	longestConsecutive := 1
	for i := 0; i < len(nums); i++ {
		mainNum := nums[i]
		count := 1

		for {
			found := false
			for j := 0; j < len(nums); j++ {
				if mainNum+1 == nums[j] {
					count++
					mainNum = nums[j]
					found = true
					break
				}
			}

			if !found {
				break
			}
		}

		if count > longestConsecutive {
			longestConsecutive = count
		}
	}

	return longestConsecutive
}

func longestConsecutiveBetter(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	sort.Ints(nums)

	longestConsecutive := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] == nums[i-1]+1 {
			count++
		} else {
			count = 1
		}
		if count > longestConsecutive {
			longestConsecutive = count
		}
	}

	return longestConsecutive
}

func longestConsecutiveOptimal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	longest := 0

	// Loop directly over the set (not nums slice)
	for num := range numsMap {
		if !numsMap[num-1] {
			currentNum := num
			count := 1

			for numsMap[currentNum+1] {
				currentNum++
				count++
			}

			if count > longest {
				longest = count
			}
		}
	}

	return longest
}

func setZeroesBetter(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowsContainsZero := make([]bool, m)
	columnContainsZero := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowsContainsZero[i] = true
				columnContainsZero[j] = true

			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowsContainsZero[i] || columnContainsZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}

//	Given an array, print all the elements which are leaders. A Leader is an element that is greater than all of the elements on its right side in the array.
//
// [9, 4, 7, 1, 0], output = [9,7,1,0]
func LeadersInArray(nums []int) []int {
	leader := nums[len(nums)-1]
	leadersArray := []int{leader}

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > leader {
			leadersArray = append(leadersArray, nums[i])
			leader = nums[i]
		}
	}

	for start, end := 0, len(leadersArray)-1; start < end; start, end = start+1, end-1 {
		leadersArray[start], leadersArray[end] = leadersArray[end], leadersArray[start]
	}

	return leadersArray
}

// If we just need to print leaders (instead of returning a slice), we can avoid extra space completely!
//
//	Output will be in reverse order, which is okay if order doesn't matter.
func LeadersInArrayOptimized(nums []int) {
	maxFromRight := nums[len(nums)-1]
	fmt.Print(maxFromRight, " ")

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= maxFromRight {
			maxFromRight = nums[i]
			fmt.Print(maxFromRight, " ")
		}
	}
}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// Traverse from Left to Right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse from Top to Bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse from Right to Left (check top <= bottom)
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse from Bottom to Top (check left <= right)
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func rotateImageBrute(matrix [][]int) [][]int {
	n := len(matrix[0])
	rotatedArr := make([][]int, n)

	for i := 0; i < n; i++ {
		rotatedArr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotatedArr[i][j] = matrix[n-1-j][i]
		}
	}

	return rotatedArr
}

func rotateImageOptimal(matrix [][]int) {
	n := len(matrix)

	// Step 1: Transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func rowWiseTraversal(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print("row:", i, " ")
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func columnWiseTraversal(matrix [][]int) {
	for j := 0; j < len(matrix[0]); j++ {
		fmt.Print("col", j, " ")
		for i := 0; i < len(matrix); i++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func reverseRowWiseTraversal(matrix [][]int) {
	for i := len(matrix) - 1; i >= 0; i-- {
		fmt.Print("row:", i, " ")
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func reverseColumnWiseTraversal(matrix [][]int) {
	for j := len(matrix) - 1; j >= 0; j-- {
		fmt.Print("reverse column wise", j, " ")
		for i := 0; i < len(matrix); i++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

// Transposing a matrix means flipping it over its diagonal — rows become columns and columns become rows.
func transposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	columns := len(matrix[0])
	result := make([][]int, columns) // because no of columns is going to be length of row for transposed matrix

	for i := 0; i < columns; i++ {
		result[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			// Transpose: element at (i, j) in original becomes (j, i) in transposed
			result[i][j] = matrix[j][i]
		}
	}

	return result
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func main() {
	// mergeBruteForce([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	// fmt.Println(countSubarrayWithSumK([]int{1, 1, 1}, 2))
	// fmt.Println(LeadersInArray([]int{9, 4, 7, 1, 0}))
	matrix := ([][]int{
		{1, 2, 3},
		{4, 5, 6},
	})
	transposedMatrix := transposeMatrix(matrix)
	printMatrix(transposedMatrix)
}
