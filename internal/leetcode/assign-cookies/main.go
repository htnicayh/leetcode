package main

import (
	"fmt"
	"sort"
)

var g = []int{1, 2, 3}
var s = []int{1, 1}

func solve(g []int, s []int) int {
	i, j := 0, 0

	sort.Ints(g)
	sort.Ints(s)

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}

		j++
	}

	return i
}

func main() {
	fmt.Println(solve(g, s))
}
