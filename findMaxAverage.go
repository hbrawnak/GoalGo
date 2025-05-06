package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	length := len(nums)
	sum := 0

	//First k elements sum
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	for i := k; i < length; i++ {
		sum -= nums[i-k] // Removing the element from very left side

		sum += nums[i] // Adding the new element from right

		maxSum = max(maxSum, sum) // finding max sum
	}

	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}
