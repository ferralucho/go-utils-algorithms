package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		current = current.Right
	}

	return result
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

	result := InOrderTraversal(root)
	fmt.Println(result)
}
