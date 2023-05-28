package tree

import "testing"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func TestValidateTree(t *testing.T) {
	// Valid binary search tree
	tree1 := &TreeNode{
		Value: 4,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 1,
			},
			Right: &TreeNode{
				Value: 3,
			},
		},
		Right: &TreeNode{
			Value: 6,
			Left: &TreeNode{
				Value: 5,
			},
			Right: &TreeNode{
				Value: 7,
			},
		},
	}
	if !validateTree(tree1, -1e9, 1e9) {
		t.Errorf("Failed to validate valid binary search tree")
	}

	// Invalid binary search tree
	tree2 := &TreeNode{
		Value: 4,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 1,
			},
			Right: &TreeNode{
				Value: 3,
			},
		},
		Right: &TreeNode{
			Value: 5,
			Left: &TreeNode{
				Value: 6,
			},
			Right: &TreeNode{
				Value: 7,
			},
		},
	}
	if validateTree(tree2, -1e9, 1e9) {
		t.Errorf("Failed to invalidate invalid binary search tree")
	}
}

func validateTree(n *TreeNode, lower, upper int) bool {
	if n == nil {
		return true
	}

	return n.Value > lower && n.Value < upper && validateTree(n.Left, lower, n.Value) && validateTree(n.Right, n.Value, upper)
}
