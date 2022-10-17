package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeStudents(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 2}, 0},
		{[]int{1, 2, 2, 3}, 0},
		{[]int{1, 5, 2, 3}, 3},
		{[]int{1, 2, 4, 5}, 0},
		{[]int{1, 5, 2, 3, 4, 5}, 4},
		{[]int{1, 2, 2, 3, 4, 5}, 0},
		{[]int{1, 6, 2, 3, 4, 5}, 5},
		{[]int{4, 5, 2, 3, 1, 6}, 5},
	}

	for _, tt := range tests {
		result := heightChecker(tt.in)
		assert.Equal(t, tt.expected, result)
	}
}
