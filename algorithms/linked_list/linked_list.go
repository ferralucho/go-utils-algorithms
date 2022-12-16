// Package linkedlist provides methods for linked lists handling.
package linkedlist

//scatalan@deuna.com

import (
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

// Insert a new node at the end of linked list
func (l *LinkedList) Insert(d interface{}) {
	list := &Node{val: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		n := l.head
		for n.next != nil {
			n = n.next
		}
		n.next = list
	}
	l.length += 1
}

// Return the number of elements in the list
func (l *LinkedList) Len() int {
	return l.length
}

// Show all values of the list
func (l *LinkedList) Show() {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.val)
		p = p.next
	}
}

// Remove node from list by value, returns the if it was removed
func (l *LinkedList) Remove(val interface{}) bool {
	var n *Node

	if l.head == nil {
		return false
	}

	if l.head.val == val {
		l.head = nil
		l.length -= 1
		return true
	}

	for n != nil && n.next != nil {
		if n.next.val == val {
			n.next = n.next.next
			l.length -= 1
			return true
		} else {
			n = n.next
		}
	}
	l.length -= 1

	return true
}

// Find a node by value in the linked list, return true if exists
func (l *LinkedList) Find(value interface{}) bool {
	var n *Node
	if l.head == nil {
		return false
	}
	n = l.head
	for n.next != nil {
		if n.val == value {
			return true
		}
		n = n.next
		if n.next == nil && n.val == value {
			return true
		}
	}
	return false
}

// Invert the linked list
func (l *LinkedList) Invert() {
	var prev *Node
	curr := l.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
