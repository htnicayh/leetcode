package main

import (
	"fmt"
	"strings"
	"unicode"
)

var s string = "2-5g-3-J7-k"
var k int = 3

func solve(s string, k int) string {
	var cleanS []rune

	for _, char := range s {
		if char != '-' {
			cleanS = append(cleanS, unicode.ToUpper(char))
		}
	}

	n := len(cleanS)
	if n == 0 {
		return ""
	}

	first := n % k

	if first == 0 {
		first = k
	}

	var res []string
	res = append(res, string(cleanS[:first]))

	for i := first; i < n; i += k {
		res = append(res, string(cleanS[i:i+k]))
	}

	return strings.Join(res, "-")
}

func main() {
	fmt.Println(solve(s, k))
}
