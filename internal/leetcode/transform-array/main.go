package main

import "fmt"

var NUMS1 []int = []int{3, -2, 1, 1}
var NUMS2 []int = []int{-1, 4, -1}
var NUMS3 []int = []int{-10, -10}

func solve(nums []int) []int {
	len := len(nums)
	ns := make([]int, len)

	for i, _ := range nums {
		index := (i + nums[i]) % len

		if index < 0 {
			index += len
		}

		ns[i] = nums[index]
	}

	fmt.Println(ns)

	return ns
}

func main() {
	solve(NUMS1)
	solve(NUMS2)
	solve(NUMS3)
}
