package main

var NUMS1 [][]int = [][]int{{1, 2}, {2, 5}, {3, 9}, {1, 15}}
var NUMS2 [][]int = [][]int{{10, 5}, {1, 7}}
var NUMS3 [][]int = [][]int{{1, 4}, {18, 5}, {15, 7}, {12, 9}, {1, 11}, {18, 13}, {16, 17}}

func solve(events [][]int) int {
	p := make(map[int]int)

	for i, vle := range events {
		index := vle[0]
		time := vle[1]

		if i == 0 {
			p[time] = index
		} else {
			sub := time - events[i-1][1]

			if _, exist := p[sub]; exist {
				if index < p[sub] {
					p[sub] = index
				}
			} else {
				p[sub] = index
			}
		}
	}

	var maxtime int
	for k, _ := range p {
		if k > maxtime {
			maxtime = k
		}
	}

	return p[maxtime]
}

func main() {
	solve(NUMS1)
	solve(NUMS2)
	solve(NUMS3)
}
