package main

import (
	"testing"
)

func TestLowestCommonAncestorBinarySearchTree(t *testing.T) {
	root := &Node{Val: 6}
	root.Left = &Node{Val: 2}
	root.Left.Left = &Node{Val: 0}
	root.Left.Right = &Node{Val: 4}
	root.Left.Right.Left = &Node{Val: 3}
	root.Left.Right.Right = &Node{Val: 5}
	root.Right = &Node{Val: 8}
	root.Right.Left = &Node{Val: 7}
	root.Right.Right = &Node{Val: 9}

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
