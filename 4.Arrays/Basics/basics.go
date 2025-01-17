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
	largest := math.MinInt
	secondLargest := math.MinInt
	for _, val := range arr {
		if val > largest {
			secondLargest = largest
			largest = val
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

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	fmt.Println(unionOfTwoArrays([]int{1, 2, 2, 3, 4}, []int{1, 1, 2, 5, 7}))
}
