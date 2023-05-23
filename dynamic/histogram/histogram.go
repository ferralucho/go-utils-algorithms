package main

import (
	"fmt"
)

func maxHistogram(input []int) int {
	stack := make([]int, 0)
	maxArea := 0
	area := 0
	i := 0

	for i < len(input) {
		if len(stack) == 0 || input[stack[len(stack)-1]] <= input[i] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				area = input[top] * i
			} else {
				area = input[top] * (i - stack[len(stack)-1] - 1)
			}

			if area > maxArea {
				maxArea = area
			}
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			area = input[top] * i
		} else {
			area = input[top] * (i - stack[len(stack)-1] - 1)
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func main() {
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
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed. Expected: %d, Got: %d\n", i+1, test.area, area)
		}
	}
}
