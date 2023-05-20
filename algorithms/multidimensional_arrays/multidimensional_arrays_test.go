package main

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {15, 18}, {8, 10}, {2, 6}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{2, 5}},
			expected:  [][]int{{2, 5}},
		},
		{
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected:  [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			intervals: [][]int{{1, 5}, {2, 7}, {6, 8}},
			expected:  [][]int{{1, 8}},
		},
	}

	for _, tc := range testCases {
		result := MergeIntervals(tc.intervals)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
