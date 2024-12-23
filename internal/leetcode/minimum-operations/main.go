package main

import "fmt"

var NUMS1 = []int{1, 2, 3, 4, 2, 3, 3, 5, 7}
var NUMS2 = []int{4, 5, 6, 4, 4}
var NUMS3 = []int{6, 7, 8, 9}

const SLICE int = 3

func solve1(nums []int) bool {
	seens := make(map[int]bool)

	for _, value := range nums {
		if seens[value] {
			return seens[value]
		}

		seens[value] = true
	}

	return false
}

func operation(nums []int) {
	n := 0
	s := solve1(nums)

	for s && len(nums) > 0 {
		n += 1

		if len(nums) < SLICE {
			nums = nums[len(nums):]
		} else {
			nums = nums[SLICE:]
		}

		s = solve1(nums)
	}

	fmt.Println(n)
}

func main() {
	operation(NUMS1)
	operation(NUMS2)
	operation(NUMS3)
}
