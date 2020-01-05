package main

import (
	"fmt"
	"sort"
)

type ElementNotFoundERROR struct {
	er string
}

func (e *ElementNotFoundERROR) Error() string {
	return e.er
}

type resp struct {
	key int
	val int
}

func BinarySearch(numbers []int, find int) (*resp, error) {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == find {
			return &resp{key: mid, val: numbers[mid]}, nil
		}

		if numbers[mid] < find {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nil, &ElementNotFoundERROR{er: "Not Found"}
}

func main() {
	numbers := []int{0, 2, 3, 5, 6, 7, 9, 10, 11, 27, 15, 30, 17, 18, 20, 23, 25}
	sort.Ints(numbers)
	fmt.Println(numbers)
	find := 30

	res, _ := BinarySearch(numbers, find)
	fmt.Println(res.key, res.val)
}
