package main

import (
	"fmt"
)

func solve(s int, t int) any {
	xor := s ^ t
	c := 0

	for xor > 0 {
		fmt.Println(xor)

		c += xor & 1
		xor >>= 1
	}

	return c
}

func main() {
	fmt.Println(solve(3, 1))
}
