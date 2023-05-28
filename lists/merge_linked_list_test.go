/*
Problem:
- Merge two sorted linked lists and return it as a new list.

Example:
- Input: 1 -> 3-> 5 & 2 -> 4-> 6
  Output: 1 -> 2-> 3 -> 4 -> 5 -> 6

Approach:
- Traverse both list at the same time, compare their values at each step and
  add the smaller one to a new list.

Solution:
- Initialize a dummy head to keep track of the head node.
- Traverse both lists while they are are not empty.
- Compare their values at each step and add the smaller one to a new list.
- Remember to check if we have reached the end for a list faster than the
  other.
- If that is the case, simply add the rest of the other list to the new list
  since it is already sorted.

Cost:
- O(n|m) time, O(n+m) space where n and m are lengths of these two linked lists.
*/

package lists

import "testing"

type ListNode struct {
	Value int
	Next  *ListNode
}

func TestMergeSortedLinkedList(t *testing.T) {
	// Test case 1
	list1 := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 3,
			Next: &ListNode{
				Value: 5,
			},
		},
	}

	list2 := &ListNode{
		Value: 2,
		Next: &ListNode{
			Value: 4,
			Next: &ListNode{
				Value: 6,
			},
		},
	}

	expectedResult := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 4,
					Next: &ListNode{
						Value: 5,
						Next: &ListNode{
							Value: 6,
						},
					},
				},
			},
		},
	}

	result := mergeSortedLinkedList(list1, list2)

	// Compare the merged list with the expected result
	for result != nil && expectedResult != nil {
		if result.Value != expectedResult.Value {
			t.Errorf("Mismatch found in merged list: Expected %d, got %d", expectedResult.Value, result.Value)
		}
		result = result.Next
		expectedResult = expectedResult.Next
	}

	if result != nil || expectedResult != nil {
		t.Errorf("Mismatch in merged list length")
	}

	// Test case 2
	list3 := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
			},
		},
	}

	list4 := &ListNode{
		Value: 4,
		Next: &ListNode{
			Value: 5,
			Next: &ListNode{
				Value: 6,
			},
		},
	}

	expectedResult2 := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 4,
					Next: &ListNode{
						Value: 5,
						Next: &ListNode{
							Value: 6,
						},
					},
				},
			},
		},
	}

	result2 := mergeSortedLinkedList(list3, list4)

	// Compare the merged list with the expected result
	for result2 != nil && expectedResult2 != nil {
		if result2.Value != expectedResult2.Value {
			t.Errorf("Mismatch found in merged list: Expected %d, got %d", expectedResult2.Value, result2.Value)
		}
		result2 = result2.Next
		expectedResult2 = expectedResult2.Next
	}

	if result2 != nil || expectedResult2 != nil {
		t.Errorf("Mismatch in merged list length")
	}
}

func mergeSortedLinkedList(l1, l2 *ListNode) *ListNode {
	mergeHead := &ListNode{}
	current := mergeHead

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return mergeHead.Next
}
