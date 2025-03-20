package main

import "fmt"

var ROOT1 = &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
var ROOT2 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode) int {
	sum := 0

	if root == nil {
		return sum
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}

	return sum + solve(root.Left) + solve(root.Right)
}

func main() {
	fmt.Println(solve(ROOT1))
	fmt.Println(solve(ROOT2))
}
