package main

import (
	"fmt"
	"slices"
)

var root []int = []int{1, 1, 0, 1, 1, 1}

func solve(nums []int) int {
	var arr []int

	count := 0

	for i, val := range nums {
		if val == 1 {
			count++
		}

		if i != len(nums)-1 && nums[i+1] == 0 || i == len(nums)-1 {
			arr = append(arr, count)

			count = 0
		}
	}

	return slices.Max(arr)
}

func main() {
	fmt.Println(solve(root))
}
