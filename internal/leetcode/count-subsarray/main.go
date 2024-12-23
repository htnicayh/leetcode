package main

import "fmt"

var NUMS1 []int = []int{1, 2, 1, 4, 1}
var NUMS2 []int = []int{1, 1, 1}
var NUMS3 []int = []int{-1, -4, -1, 4}

func solve(nums []int) bool {
	sum := nums[0] + nums[2]
	d := nums[1] / 2
	m := nums[1] % 2

	return m == 0 && sum == d
}

func countSubarrays(nums []int) int {
	if len(nums) == 3 {
		b := solve(nums)

		if b {
			return 1
		}

		return 0
	}

	n := 0

	for i, _ := range nums {
		if i < len(nums)-2 && solve(nums[i:i+3]) {
			n = n + 1
		}
	}

	fmt.Println(n)

	return n
}

func main() {
	countSubarrays(NUMS1)
	countSubarrays(NUMS2)
	countSubarrays(NUMS3)
}
