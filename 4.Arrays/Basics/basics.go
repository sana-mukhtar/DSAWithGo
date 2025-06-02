package main

import (
	"fmt"
	"math"
)

func findLargest(arr []int) int {
	max := math.MinInt
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func findSecondLargest(arr []int) int {
	largest, secondLargest := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		}
		if arr[i] > secondLargest && arr[i] < largest {
			secondLargest = arr[i]
		}
	}

	return secondLargest
}

func isSortedAsc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func isSortedDesc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

func isSorted(arr []int) bool {
	if arr[0] > arr[1] {
		return isSortedDesc(arr)
	} else {
		return isSortedAsc(arr)
	}
}

func removeDuplicates(nums []int) int {
	var uniqueArray []int
	uniqueArray = append(uniqueArray, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		uniqueArray = append(uniqueArray, nums[i])
	}
	fmt.Println("unique array: ", uniqueArray)
	return len(uniqueArray)
}

// [0,1,1,2,3,4,5,5] array wil be in increasing order
// if i and j are not same continue, return the length of unique elements
func removeDuplicatesUsingTwoPointers(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func rotate(nums []int, k int) {
	length := len(nums) - 1  // 6
	for i := 0; i < k; i++ { // 3 steps
		temp := nums[length]
		for j := length; j > 0; j-- {
			fmt.Println(nums[j], ":", nums[j-1])
			nums[j] = nums[j-1]
		}
		nums[0] = temp
		fmt.Println(i, ":", nums)
	}
}

// func moveZeroes(nums []int) {
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] == 0 {
// 			temp := nums[i]
// 			for j := i; j < len(nums)-1; j++ {
// 				nums[j] = nums[j+1]
// 			}
// 			nums[len(nums)-1] = temp
// 			fmt.Println(i, ":",nums)
// 		} else {
// 			continue
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	nonZeroIdx := 0
	for _, num := range nums {
		if num != 0 {
			nums[nonZeroIdx] = num
			nonZeroIdx++
		}
	}

	for i := nonZeroIdx; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroesUsingTwoPointers(arr []int) []int {
	zeroVal := 0
	for i := 0; i < len(arr); i++ {
		// swap with zero value element when arr[i] is non zero
		if arr[i] != 0 {
			arr[zeroVal], arr[i] = arr[i], arr[zeroVal]
			zeroVal++
		}
	}
	return arr
}

func linearSearch(key int, arr []int) int {
	for i, v := range arr {
		if v == key {
			return i
		}
	}
	return -1
}

func unionOfTwoArrays(arr1, arr2 []int) []int {
	m, n := len(arr1), len(arr2)
	i, j := 0, 0
	union := []int{}

	for i < m && j < n {
		// equal condition
		if arr1[i] == arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
			j++
		} else if arr1[i] < arr2[j] { // small condition
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else { //greater condition
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		}
	}

	for i < m {
		if len(union) == 0 || union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}
	for j < n {
		if len(union) == 0 || union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}
	return union
}

func interSectionOfTwoArrays(arr1, arr2 []int) []int {
	m, n := len(arr1), len(arr2)
	i, j := 0, 0
	intersectionArray := []int{}
	for i < m && j < n {
		if arr1[i] == arr2[j] {
			intersectionArray = append(intersectionArray, arr1[i])
			i++
			j++
		}
		if arr1[i] > arr2[j] {

		}
	}
	return intersectionArray
}

func missingNumber(nums []int) int {
	for i := 0; i <= len(nums); i++ {
		flag := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return i
		}
	}
	return -1
}

// better
func missingNumberBetter(nums []int) int {
	sum1, sum2 := 0, 0
	for i := 0; i <= len(nums); i++ {
		sum1 += i
	}

	for i := 0; i < len(nums); i++ {
		sum2 += nums[i]
	}
	return sum1 - sum2
}

// optimal
func missingNumberOptimal(nums []int) int {
	sum1, sum2 := 0, 0
	length := len(nums)
	sum1 = length * (length + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum2 += nums[i]
	}
	return sum1 - sum2
}

func findMaxConsecutiveOnes(nums []int) int {
	max, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if max < count {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == nums[i] {
				count++
			}
		}
		if count != 2 {
			return nums[i]
		}
	}
	return -1
}

func subArrayWithSumK(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

func longestsubArrayWithSumK(arr []int, k int) []int {
	longestSubArray := []int{}
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			subArrLen := j - i + 1
			if sum == k && subArrLen > len(longestSubArray) {
				longestSubArray = arr[i : j+1]
			}
		}
	}
	return longestSubArray
}

// [1,2,3,4,5]
func rotateByOnePlace(arr []int) []int {
	length := len(arr)
	element := arr[length-1]
	for i := length - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = element
	return arr
}

// [1,2,3,4,5,6,7,8] n=3 [6,7,8] len = 8-1-3
// append 8-3
func rotateByNPlaces(arr []int, n int) []int {
	rotatedArr := arr[len(arr)-n:]
	rotatedArr = append(rotatedArr, arr[:len(arr)-1-n]...)
	return rotatedArr
}

// Q: Check if a pair exists with a given sum in a sorted array
func checkSum(arr []int, target int) bool {
	left := arr[0]
	right := arr[len(arr)-1]
	for left < right {
		if left+right > target {
			right--
		} else if left+right < target {
			left++
		} else {
			return true
		}
	}

	return false
}

func main() {
	// fmt.Println(findSecondLargest([]int{2, 6, 9, 4, 2, 0, 1, 8}))
	// fmt.Println(checkSum([]int{1, 2, 4, 5, 6}, 13))
	// fmt.Println(longestsubArrayWithSumK([]int{10, 2, 5, 2, 5, 20, 1, 2, 4}, 10))
	fmt.Println(moveZeroesUsingTwoPointers([]int{0, 1, 0, 3, 12}))
}
