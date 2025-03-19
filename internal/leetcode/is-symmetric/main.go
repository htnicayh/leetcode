package main

import "fmt"

var ROOT = &TreeNode{1,
	&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
	&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
}

func solve(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)
}

func main() {
	fmt.Println(solve(ROOT))
}
