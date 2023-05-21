package main

import "fmt"

// TreeNode represents a node in the binary tree
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
		//Traverse to the left most node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		//Process the current node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		//Traverse to the right tree
		current = current.Right
	}

	return result
}

func main() {
	// Create a sample binary tree
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
