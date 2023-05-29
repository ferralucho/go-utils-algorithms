/*
Problem:
- Given a list of movie lengths (integer) and a flight length (integer), determine if
  there exist two movies that add up to the total length. Assume that an user
  watch exactly two movies, but not the same one twice.

Example:
- Input: list=[]int{2, 3, 4}, length=6
  Output: true, because there exists 2 and 4 that add up to 6

Approach:
- Could use hashmap to keep track of movie lengths that we've seen so far and
  look it up as we iterate through the list.

Solution:
- Use a hashmap to keep track of movie lengths that we've seen so far.
- As we iterate through the list, we calculate the difference for each value
  (total length - current movie length) and check if the difference is equal
  to the movie length that we've seen.
- Return true if that's the case.

Cost:
- O(n) time, O(n) space.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFillFlights(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected bool
	}{
		{[]int{}, 1, false},
		{[]int{0}, 1, false},
		{[]int{1}, 1, false},
		{[]int{0, 1}, 1, true},
		{[]int{1, 1}, 2, true},
		{[]int{2, 3, 4}, 6, true},
		{[]int{2, 3, 4}, 8, false},
	}

	for _, tt := range tests {
		result := fillFlight(tt.in1, tt.in2)
		assert.Equal(t, tt.expected, result)
	}
}

func fillFlight(moviesLength []int, fl int) bool {
	movies := map[int]int{}

	for _, v := range moviesLength {
		matchLength := fl - v

		if _, ok := movies[matchLength]; ok {
			return true
		}
		movies[v] = 1
	}
	return false
}
