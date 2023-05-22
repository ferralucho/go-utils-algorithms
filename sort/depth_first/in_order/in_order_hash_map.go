package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	hashMap := make(map[*TreeNode]bool)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if node.Left != nil && !hashMap[node.Left] {
			stack = append(stack, node.Left)
		} else {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			hashMap[node] = true

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
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
