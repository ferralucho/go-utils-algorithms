package main

import (
	"testing"
)

func TestMaxHistogram(t *testing.T) {
	tests := []struct {
		input []int
		area  int
	}{
		{input: []int{2, 2, 2, 6, 1, 5, 4, 2, 2, 2, 2}, area: 12},
		{input: []int{1, 2, 3, 4, 5}, area: 9},
		{input: []int{5, 4, 3, 2, 1}, area: 9},
		{input: []int{2, 2, 2, 2, 2}, area: 10},
		{input: []int{1}, area: 1},
		{input: []int{0}, area: 0},
		{input: []int{}, area: 0},
	}

	for i, test := range tests {
		area := maxHistogram(test.input)
		if area == test.area {
			t.Logf("Test case %d passed", i+1)
		} else {
			t.Errorf("Test case %d failed. Expected: %d, Got: %d", i+1, test.area, area)
		}
	}
}
