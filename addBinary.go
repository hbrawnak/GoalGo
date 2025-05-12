package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		bitA := 0
		bitB := 0

		if i >= 0 {
			bitA = int(a[i] - '0')
			i--
		}

		if j >= 0 {
			bitB = int(b[j] - '0')
			j--
		}

		total := bitA + bitB + carry
		result = string('0'+(total%2)) + result
		carry = total / 2
	}

	return result
}

func main() {
	fmt.Println(addBinary("1011", "1101")) // Output: "11000"
}
