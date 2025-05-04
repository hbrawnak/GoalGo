package main

import "fmt"

func findMax(nums []int, n int) int {
	if n == 0 {
		panic("empty nums")
	}

	if n == 1 {
		return nums[0]
	}

	maxValue := nums[0]
	for _, num := range nums[1:] {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}

func main() {
	nums := []int{2, 3}
	maxNum := findMax(nums, len(nums))
	fmt.Println(maxNum)
}
