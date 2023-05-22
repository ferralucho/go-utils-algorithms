package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func PreOrderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	preOrderTraversal(root, &result)
	return result
}

func preOrderTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)
	preOrderTraversal(node.Left, result)
	preOrderTraversal(node.Right, result)
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

	result := PreOrderTraversal(root)
	fmt.Println(result)
}
