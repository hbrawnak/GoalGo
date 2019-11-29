package main

import "fmt"

type ElementNotFoundERROR struct {
	er string
}

func (e *ElementNotFoundERROR) Error() string  {
	return e.er
}

func BinarySearch()  {
	
}


func main() {
	fmt.Println("Binary")
}