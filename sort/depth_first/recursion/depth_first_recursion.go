package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DepthFirstSearch(root *TreeNode) []int {
	result := []int{}
	dfs(root, &result)
	return result
}

func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)

	dfs(node.Left, result)
	dfs(node.Right, result)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := DepthFirstSearch(root)

	fmt.Println(result)
}
