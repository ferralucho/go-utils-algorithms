package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 9, 2, 6, 1, 8, 3, 7},
			expected: []int{1, 2, 3, 5, 6, 7, 8, 9},
		},
		{
			input:    []int{4, 3, 2, 1},
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		inputCopy := make([]int, len(tc.input))
		copy(inputCopy, tc.input)

		MergeSort(inputCopy)

		if !reflect.DeepEqual(inputCopy, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, inputCopy)
		}
	}
}
