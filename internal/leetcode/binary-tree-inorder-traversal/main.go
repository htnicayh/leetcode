package main

import "fmt"

var ROOT = &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode) []int {
	var res []int

	inorder(root, &res)

	return res
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}

func main() {
	fmt.Println(solve(ROOT))
}
