package main

import (
	"testing"
)

func TestSearchProductsBinary(t *testing.T) {
	productIDs := []int{1, 3, 5, 7, 9, 11, 13, 15}
	encontrar := 9

	index, iterations := searchProductsBinary(productIDs, encontrar)

	expectedIndex := 4
	expectedIterations := 3

	if index != expectedIndex {
		t.Errorf("Expected index %d, but got %d", expectedIndex, index)
	}

	if iterations != expectedIterations {
		t.Errorf("Expected iterations %d, but got %d", expectedIterations, iterations)
	}
}
