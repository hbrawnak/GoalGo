package main

import "fmt"

// https://leetcode.com/problems/repeated-dna-sequences/

func findRepeatedDnaSequences(s string) []string {
	seen := make(map[string]bool)
	repeated := make(map[string]bool)
	var result []string

	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]
		if seen[sub] {
			repeated[sub] = true
		} else {
			seen[sub] = true
		}
	}

	for seq := range repeated {
		result = append(result, seq)
	}

	return result
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	//s := "AAAAAAAAAAAAA"
	//s := "AAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
