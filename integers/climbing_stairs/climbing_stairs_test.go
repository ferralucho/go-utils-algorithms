package main

import (
	"fmt"
	"testing"
)

/*
Problem:
- You are climbing a staircase. It takes n steps to reach to the top.
- Each time you can either climb 1 or 2 steps. In how many distinct ways
  can you climb to the top.

Example:
- Input: 2
  Output: 2
  Explanation: Two ways to climb the top are 1+1 and 2.
- Input: 3
  Output: 3
  Explanation: Two ways to climb the top are 1+1+1, 1+2, and 2+1.

Approach:
- To reach n step, one either advance one step from the n-1 step or
  2 steps from the n-2 step.
- This is basically the same recurrence formula defined by the Fibonacci
  sequence.

Solution:
- Initialize two variables to store 2 previous value.
- Iterate from 2 to n and update these two values at each step.

Cost:
- O(n) time, O(1) space.
*/

func TestClimbStairs(t *testing.T) {
	// Test case 1
	n1 := 2
	expectedResult1 := 2
	result1 := climbStairs(n1)
	if result1 != expectedResult1 {
		t.Errorf("Mismatch found: Expected %d, got %d", expectedResult1, result1)
	}

	// Test case 2
	n2 := 3
	expectedResult2 := 3
	result2 := climbStairs(n2)
	if result2 != expectedResult2 {
		t.Errorf("Mismatch found: Expected %d, got %d", expectedResult2, result2)
	}

	// Test case 3
	n3 := 4
	expectedResult3 := 5
	result3 := climbStairs(n3)
	if result3 != expectedResult3 {
		t.Errorf("Mismatch found: Expected %d, got %d", expectedResult3, result3)
	}
}

func climbStairs(n int) int {
	p, q := 1, 1

	for i := 2; i <= n; i++ {
		tmp := q
		q += p
		p = tmp
	}
	fmt.Println("q: ", q)
	return q
}
