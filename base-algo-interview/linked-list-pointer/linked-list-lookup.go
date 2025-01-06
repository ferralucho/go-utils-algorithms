/*
Fast and Slow Pointers for Linked Lists
- Use two pointers moving at different speeds to detect cycles or find midpoints.
- This approach avoids extra memory usage and works in O(n) time.
- Example: Detect if a linked list has a loop.
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func findMiddle(head *ListNode) *ListNode {
	if hasCycle(head) {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	// Create a sample linked list with a cycle
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	// Create a cycle
	// head.Next.Next.Next.Next = head.Next

	if hasCycle(head) {
		fmt.Println("Cycle detected in the linked list")
	} else {
		fmt.Println("No cycle detected in the linked list")
	}

	middle := findMiddle(head)
	fmt.Print("middle: ", middle.Val)
}
