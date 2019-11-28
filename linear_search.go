package main

import "fmt"

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

func search(numbers []int, find int) (*resp, error) {
	for index, value := range numbers {
		if find == value {
			return &resp{key: index, val: value}, nil
		}
	}
	return nil, &ElementNotFoundERROR{er: "Not Found"}
}

func main() {
	numbers := []int{3, 4, 5, 6, 7, 9, 0, 3, 11, 23}
	find := 11
	found, err := search(numbers, find)
	if err != nil {
		fmt.Println("Not Found")
	} else {
		fmt.Println(found.key, found.val)
	}
}
