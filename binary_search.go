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

/*type resp struct {
	val int
}
*/
func BinarySearch(numbers []int, find int) int {

	low := 0
	high := len(numbers) - 1
	mid := int(low + (high-low)/2)

	for low < high {
		if numbers[mid] > find {
			high = mid
		} else if numbers[mid] < find {
			low = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	numbers := []int{0, 2, 3, 5, 6, 7, 9, 10, 11, 27, 15, 30, 17, 18, 20, 23, 25}
	sort.Ints(numbers)
	find := 18

	fmt.Println(BinarySearch(numbers, find))
}
