package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

/**
 * Video link - https://youtu.be/D2bIbWGgvzI
 *
 * Given a binary tree print its level order traversal in reverse
 * e.g           1
 *          2         3
 *        4    5    6   7
 *
 * Output should be 4 5 6 7 2 3 1
 *
 * Solution
 * Maintain a stack and queue. Do regular level order traversal but
 * put right first in the queue. Instead of printing put the result
 * in stack. Finally print contents of the stack.
 *
 * Time and space complexity is O(n)
 *
 * References : http://www.geeksforgeeks.org/reverse-level-order-traversal/
 */

func reverseLevelOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	queue := list.New()
	stack := list.New()

	queue.PushBack(root)
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(*Node)

		if node.right != nil {
			queue.PushBack(node.right)
		}
		if node.left != nil {
			queue.PushBack(node.left)
		}

		stack.PushBack(node)
	}

	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		node := element.Value.(*Node)

		fmt.Printf("%d ", node.data)
	}
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	fmt.Println("Reverse Level Order Traversal:")
	reverseLevelOrderTraversal(root)
}
