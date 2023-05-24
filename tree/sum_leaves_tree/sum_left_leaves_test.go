package main

import "testing"

func createBinaryTree(nums []int) *Tree {
	return performCreate(nums, 0)
}

func performCreate(nums []int, index int) *Tree {
	if index >= len(nums) {
		return nil
	}

	tree := Tree{Val: nums[index]}
	tree.Left = performCreate(nums, 2*index+1)
	tree.Right = performCreate(nums, 2*index+2)
	return &tree
}

func TestSumOfLeftLeaves(t *testing.T) {
	root := createBinaryTree([]int{3, 9, 20, 15, 7, 2})
	expected := 17

	if res := sumOfLeftLeaves(root); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
