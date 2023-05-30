/*
Problem:
- Determine whether an integer is a palindrome.

Assumption:
- Do this without extra space.
- Define negative integers as non-palindrome.

Example:
- Input: 101
  Output: true
- Input: 106
  Output: false

Approach:
- Use two-pointer approach where one starts at the first digit and one starts
  at the last digit, have them walk toward the middle and compare them at each
  step.

Solution:
- Calculate the division factor for the number to divide by to get its first digit.
- While the number is not equal to 0:
  - Get the first digit by diving the number by the division factor above.
  - Get the last digit by applying modulo the number by 10.
- Return false immediately if they are equal.
- Otherwise, chop of both digit by applying modulo the number by the division factor
  divide the result by 10.
- Update the division factor by dividing it by 100 since we already chopped of the first
  and last digits.

Cost:
- O(n) time, O(1) space.
*/

package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIntPalindrome(t *testing.T) {
	tests := []struct {
		in       int
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, true},
		{10, false},
		{11, true},
		{101, true},
		{111, true},
		{110, false},
		{-101, false},
	}

	for _, tt := range tests {
		result := isNumberPalindrome(tt.in)
		assert.Equal(t, tt.expected, result)
	}
}

func isNumberPalindrome(x int) bool {
	str := strconv.Itoa(x)

	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}

	return true
}
