package linkedlist

import (
	"testing"
)

func createSinglyLinkedList(nums []int) *LinkedList {
	var cur *Node
	list := &LinkedList{}

	for i, num := range nums {
		if i == 0 {
			cur = &Node{val: num}
			list.head = cur
		} else {
			cur.next = &Node{val: num}
			cur = cur.next
		}
	}

	list.length = len(nums)
	return list
}

func TestInsertNodes(t *testing.T) {
	list := createSinglyLinkedList([]int{1, 2, 6, 3, 4, 5})
	val := 8
	expected := createSinglyLinkedList([]int{1, 2, 3, 4, 5, 8})

	list.Insert(val)
	if list.length == expected.length {
		t.Errorf("expected %v, got %v", expected.length, list.length)
	}
}

func TestRemoveNodes(t *testing.T) {
	list := createSinglyLinkedList([]int{1, 2, 6, 3, 4, 5})
	val := 5
	expected := createSinglyLinkedList([]int{1, 2, 3, 4})

	if res := list.Remove(val); res && list.length == expected.length {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestRemoveMultipleNodes(t *testing.T) {
	list := createSinglyLinkedList([]int{1, 2, 6, 3, 4, 5})
	firstVal := 5
	expected := createSinglyLinkedList([]int{1, 2, 3, 4})

	if res := list.Remove(firstVal); res && list.length == expected.length {
		t.Errorf("expected %v, got %v", true, res)
	}

	secondVal := 2
	secondExpected := createSinglyLinkedList([]int{1, 3, 4})

	if res := list.Remove(secondVal); res && list.length == secondExpected.length {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestFindByExistentValue(t *testing.T) {
	list := createSinglyLinkedList([]int{1, 2, 6, 3, 4, 5})
	val := 5

	if res := list.Find(val); !res {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestFindByNonExistentValue(t *testing.T) {
	list := createSinglyLinkedList([]int{1, 2, 6, 3, 4, 5})
	val := 11

	if res := list.Find(val); res {
		t.Errorf("expected %v, got %v", true, res)
	}
}
