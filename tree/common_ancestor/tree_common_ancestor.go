package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

/**
 * Lowest common ancestor in binary search tree.
 *
 * Time complexity O(height of tree)
 * Space complexity O(height of tree)
 *
 * https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
 */

func LowestCommonAncestorBinarySearchTree(root *Node, p, q int) *Node {
	fp := float64(p)
	fq := float64(q)
	if root.Val > int(math.Max(fp, fq)) && root.Left != nil {
		return LowestCommonAncestorBinarySearchTree(root.Left, p, q)
	} else if root.Val < int(math.Min(fp, fq)) && root.Right != nil {
		return LowestCommonAncestorBinarySearchTree(root.Right, p, q)
	} else {
		return root
	}
}

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}

	ca := LowestCommonAncestorBinarySearchTree(root, 6, 7)
	fmt.Println(ca.Val)
}
