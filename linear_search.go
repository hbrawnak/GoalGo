package main

import "fmt"

func search(numbers []int, find int) (x, y int) {
	key := 0
	for index, val := range numbers {
		if find == val {
			return index, val
		}
	}
	return key, key
}

func main() {
	numbers := []int{3, 4, 5, 6, 7, 9, 0, 3, 11, 23}
	find := 7
	fmt.Println(search(numbers, find))
}
