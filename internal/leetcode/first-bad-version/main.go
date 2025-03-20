package main

import "fmt"

const bad int = 4

func isBad(version int) bool {
	return version >= bad
}

func solve(n int) int {
	left, right := 1, n

	for left < right {
		middle := (left + right) / 2

		if isBad(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}

func main() {
	fmt.Println(solve(5))
}
