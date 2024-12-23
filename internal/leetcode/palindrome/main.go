package main

import (
	"fmt"
	"strconv"
)

func solve(x int) bool {
	s := strconv.Itoa(x)

	for j := 0; j < len(s)/2; j++ {
		if s[j] != s[len(s)-1-j] {
			return false
		}
	}

	return true
}

func palindrome() {
	a := solve(121)

	fmt.Println(a)
}

func main() {
	palindrome()
}
