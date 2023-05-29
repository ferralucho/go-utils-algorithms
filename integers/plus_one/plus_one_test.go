package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Problem:
- Given a number represented as an array of digits, plus one to the number.

Assumption:
- The input are non-negative.
- The digits are stored such that the most significant digit is at the head of the list.
- The number does not contain leading zeros.

Example:
- Input: []int{1, 2, 5}
  Output: []int{1, 2, 6}
- Input: []int{1, 2, 9}
  Output: []int{1, 3, 0}
- Input: []int{1, 9, 9}
  Output: []int{2, 0, 0}

Solution:
- Iterate through the list from right to left and add 1 to the current digit accordingly.
  - If the current digit is less than 9, add 1 and update it.
  - Otherwise, set it to 0.
- If all the digits are 9, append an 0 in the end and update the first digit to 1.

Cost:
- O(n) time, O(1) space, where n is the length of the list.
*/

func TestPlusOne(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{1}},
		{[]int{1}, []int{2}},
		{[]int{9}, []int{1, 0}},
		{[]int{1, 2, 5}, []int{1, 2, 6}},
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}

	for _, tt := range tests {
		result := plusOne(tt.in)
		assert.Equal(t, tt.expected, result)
	}
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	counter := 0

	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			break
		} else {
			digits[i] = 0
			counter++
		}
	}

	if counter == len(digits) {
		digits = append(digits, 0)
		digits[0] = 1
	}

	return digits
}
