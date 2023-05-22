package main

import (
	"reflect"
	"testing"
)

func TestInOrderTraversal(t *testing.T) {
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

	expected := []int{4, 2, 5, 1, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
