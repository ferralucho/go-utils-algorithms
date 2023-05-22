package main

import (
	"reflect"
	"testing"
)

func TestInOrderTraversal(t *testing.T) {
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

	// Perform in-order traversal and get the result
	result := InOrderTraversal(root)

	// Define the expected result
	expected := []int{4, 2, 5, 1, 3}

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
