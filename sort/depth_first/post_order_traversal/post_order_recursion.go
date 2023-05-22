package main

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (root *Node) PostOrderTraversal() []int {
	result := make([]int, 0)
	postOrder(root, &result)
	return result
}

func postOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}

	postOrder(node.left, result)
	postOrder(node.right, result)
	*result = append(*result, node.data)
}
