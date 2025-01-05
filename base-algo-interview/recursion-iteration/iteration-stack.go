/*
	Convert Recursion to Iteration with a Stack
- Recursive algorithms can be converted to iterative ones using a stack.
- This approach provides more control over memory and avoids stack overflow.
- Example: Iterative in-order traversal of a binary tree.
*/

package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		// Traverse to the leftmost node
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// Pop the top node from the stack and visit it
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)

		// Move to the right subtree
		root = root.Right
	}

	return result
}

func main() {
	// Example binary tree
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	result := inorderTraversal(root)
	fmt.Println("In-order traversal:", result) // Output: In-order traversal: [4 2 5 1 3]
}
