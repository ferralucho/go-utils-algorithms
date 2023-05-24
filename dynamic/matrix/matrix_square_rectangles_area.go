package main

import "fmt"

/*
Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

Example 1:

Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 6
Explanation: The maximal rectangle is shown in the above picture.
Example 2:

Input: matrix = [["0"]]
Output: 0
Example 3:

Input: matrix = [["1"]]
Output: 1


Constraints:

rows == matrix.length
cols == matrix[i].length
1 <= row, cols <= 200
matrix[i][j] is '0' or '1'.

 * Given a 2D matrix of 0s and 1s. Find largest rectangle of all 1s
 * in this matrix.
 *
 * Maintain a temp array of same size as number of columns.
 * Copy first row to this temp array and find largest rectangular area
 * for histogram. Then keep adding elements of next row to this temp
 * array if they are not zero. If they are zero then put zero there.
 * Every time calculate max area in histogram.
 *
 * Time complexity - O(rows*cols)
 * Space complexity - O(cols) - if number of cols is way higher than rows
 * then do this process for rows and not columns.
 *
 * References:
 * http://www.careercup.com/question?id=6299074475065344

*/

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

func maximumRectangularSubmatrixOf1s(input [][]int) int {
	temp := make([]int, len(input[0]))
	maxArea := 0
	area := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(temp); j++ {
			if input[i][j] == 0 {
				temp[j] = 0
			} else {
				temp[j] += input[i][j]
			}
		}
		area = maxHistogram(temp)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	input := [][]int{
		{1, 1, 1, 0},
		{1, 1, 1, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1},
	}
	maxRectangle := maximumRectangularSubmatrixOf1s(input)
	fmt.Println("Max rectangle is of size", maxRectangle)
}
