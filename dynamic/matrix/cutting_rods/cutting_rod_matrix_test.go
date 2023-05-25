package main

import (
	"testing"
)

func TestMaxValue(t *testing.T) {
	prices := []int{3, 5, 8, 9, 10, 20, 22, 25}
	expected := 26
	result := MaxValue(prices)
	if result != expected {
		t.Errorf("MaxValue(%v) = %d; want %d", prices, result, expected)
	}
}
