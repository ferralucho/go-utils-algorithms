package main

import "math"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func LowestCommonAncestorBinarySearchTree(root *TreeNode, p, q int) *TreeNode {
	fp := float64(p)
	fq := float64(q)
	if root.Val > int(math.Max(fp, fq)) {
		return LowestCommonAncestorBinarySearchTree(root.Left, p, q)
	} else if root.Val < int(math.Min(fp, fq)) {
		return LowestCommonAncestorBinarySearchTree(root.Right, p, q)
	} else {
		return root
	}
}
