package main

/*
404. Sum of Left Leaves
https://leetcode.com/problems/sum-of-left-leaves/

Find the sum of all left leaves in a given binary tree.
*/

type Tree struct {
	Val   int
	Right *Tree
	Left  *Tree
}

// time complexity: O(n), where n is nodes number in the tree.
// space complexity: O(h), where h is height of the tree.
func performSum(root *Tree, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	return performSum(root.Left, true) + performSum(root.Right, false)
}

func sumOfLeftLeaves(root *Tree) int {
	return performSum(root, false)
}
