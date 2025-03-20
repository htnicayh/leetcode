package main

import "fmt"

var NUMS1 = []int{3, 2, 3}
var NUMS2 = []int{2, 2, 1, 1, 1, 2, 2}

func solve(nums []int) int {
	target, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			target = num
		}

		if num == target {
			count++
		} else {
			count--
		}
	}

	return target
}

func main() {
	fmt.Println(solve(NUMS1))
	fmt.Println(solve(NUMS2))
}
