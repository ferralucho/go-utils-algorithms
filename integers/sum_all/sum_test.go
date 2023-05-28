package main

import (
	"reflect"
	"testing"
)

func TestFindSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{nums: []int{1, 2, 3}, target: 4, expected: []int{0, 2}},
		{nums: []int{}, target: 10, expected: []int{0, 0}},
		{nums: []int{5}, target: 5, expected: []int{0, 0}},
	}

	for _, test := range tests {
		result := FindSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Failed for input nums=%v, target=%d. Expected: %v, got: %v",
				test.nums, test.target, test.expected, result)
		}
	}
}
