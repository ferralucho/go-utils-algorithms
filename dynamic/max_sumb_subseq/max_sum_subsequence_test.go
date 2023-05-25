package main

import "testing"

func TestMaxSum(t *testing.T) {
	arr1 := []int{1, 101, 10, 2, 3, 100, 4}
	expected1 := 111
	result1 := MaxSum(arr1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed. Expected: %d, got: %d", expected1, result1)
	}

	arr2 := []int{3, 4, 5, 10}
	expected2 := 22
	result2 := MaxSum(arr2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed. Expected: %d, got: %d", expected2, result2)
	}

	arr3 := []int{10, 5, 4, 3}
	expected3 := 10
	result3 := MaxSum(arr3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed. Expected: %d, got: %d", expected3, result3)
	}
}
