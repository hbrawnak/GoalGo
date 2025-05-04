package main

import (
	"fmt"
	"math"
)

func findSecondLargest(nums []int) int {
	highest := nums[0]
	secondHighest := math.MinInt32

	for i := 1; i < len(nums); i++ {
		if nums[i] > highest {
			secondHighest = highest
			highest = nums[i]
		} else if nums[i] < highest && nums[i] > secondHighest {
			secondHighest = nums[i]
		}
	}

	if secondHighest == math.MinInt32 {
		return -1
	}

	return secondHighest
}

func main() {
	arr := []int{10, 5, 10}
	//arr := []int{10, 10, 10}
	//arr := []int{12, 35, 1, 10, 34, 1}
	secondLargest := findSecondLargest(arr)
	fmt.Println(secondLargest)
}

// T: O(n) S: O(1)
