package main

import "testing"

func TestMaximumSizeSubMatrix(t *testing.T) {
	arr := [][]int{
		{0, 1, 1, 0, 1},
		{1, 1, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 0, 1},
	}

	expected := 3
	result := MaximumSizeSubMatrix(arr)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
