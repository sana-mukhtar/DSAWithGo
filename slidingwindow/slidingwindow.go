package main

import "fmt"

func maxSubArraySum(arr []int, k int)int{
    n := len(arr)
    
    if n <k {
        return -1
    }
    
    windowSum:=0
    for i :=0; i<k; i++{
        windowSum += arr[i]
    }
    
    maxSum := windowSum
	for i := k; i < n; i++ {
        windowSum = windowSum-arr[i-k]+arr[i]
        
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }
    return maxSum
}

func main() {
  fmt.Println("Start small. Ship something.", maxSubArraySum([]int{2, 1, 5, 1, 3, 2}, 3))
}