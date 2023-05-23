package main

import (
	"testing"
)

func TestLowestCommonAncestorBinarySearchTree(t *testing.T) {
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}

	tests := []struct {
		p, q   int
		expect int
	}{
		{2, 8, 6},
		{2, 4, 2},
		{3, 7, 6},
	}

	for _, test := range tests {
		result := LowestCommonAncestorBinarySearchTree(root, test.p, test.q)
		if result.Val != test.expect {
			t.Errorf("LowestCommonAncestorBinarySearchTree(%d, %d) returned %d, expected %d",
				test.p, test.q, result.Val, test.expect)
		}
	}
}
