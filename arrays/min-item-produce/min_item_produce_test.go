package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinItemProduce(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 2}, 0},
		{[]int{1, 2, 2, 3}, 4},
		{[]int{1, 5, 2, 3}, 5},
		{[]int{1, 2, 4, 5}, 4},
		{[]int{1, 5, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 2, 3, 4, 5}, 4},
		{[]int{1, 6, 2, 3, 4, 5}, 5},
		{[]int{4, 5, 2, 3, 1, 6}, 8},
	}

	for _, tt := range tests {
		result := MinItemProduce(3, 5, tt.in)
		assert.Equal(t, tt.expected, result)
	}
}
