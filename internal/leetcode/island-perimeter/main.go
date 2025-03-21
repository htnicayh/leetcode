package main

import "fmt"

var root [][]int = [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}

func solve(grid [][]int) int {
	is := 0

	for i := 0; i < len(grid); i++ {
		row := grid[i]

		for j := 0; j < len(row); j++ {
			cell := row[j]

			if cell == 1 {
				is += 4

				if j != len(row)-1 && cell == row[j+1] {
					is -= 2
				}

				if i != len(grid)-1 && cell == grid[i+1][j] {
					is -= 2
				}
			}
		}
	}

	return is
}

func main() {
	fmt.Println(solve(root))
}
