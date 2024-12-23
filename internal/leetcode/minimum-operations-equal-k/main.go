package main

var NUMS1 []int = []int{5, 2, 5, 4, 5}
var NUMS2 []int = []int{2, 1, 2}
var NUMS3 []int = []int{9, 7, 5, 3}

const K1 int = 2
const K2 int = 2
const K3 int = 1

func quickSortInPlace(nums []int, l int, r int) []int {
	if l < r {
		pivot := partition(nums, l, r)

		quickSortInPlace(nums, l, pivot-1)
		quickSortInPlace(nums, pivot+1, r)
	}

	return nums
}

func partition(nums []int, l int, r int) int {
	pivot := nums[r]
	i := l - 1

	for j := l; j < r; j++ {
		if nums[j] > pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[r] = nums[r], nums[i+1]

	return i + 1
}

func solve(nums []int, k int) int {
	nums = quickSortInPlace(nums, 0, len(nums)-1)

	if nums[len(nums)-1] < k {
		return -1
	}

	n := 0

	for j := 0; j < len(nums); j++ {
		if j < len(nums)-1 {
			if nums[j] != nums[j+1] {
				n++
			}
		} else if nums[j] > k {
			n++
		}
	}

	return n
}

func main() {
	solve(NUMS1, K1)
	solve(NUMS2, K2)
	solve(NUMS3, K3)
}
