package main

import "fmt"

/**
 *
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

// MaximumSizeSubMatrix calculates the maximum size of a submatrix with all 1s in a binary matrix.
func MaximumSizeSubMatrix(arr [][]int) int {
	m := len(arr)
	n := len(arr[0])

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	maxSize := 0
	for i := 0; i < m; i++ {
		result[i][0] = arr[i][0]
		if result[i][0] == 1 {
			maxSize = 1
		}
	}

	for i := 0; i < n; i++ {
		result[0][i] = arr[0][i]
		if result[0][i] == 1 {
			maxSize = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if arr[i][j] == 0 {
				continue
			}
			t := min(result[i-1][j], result[i-1][j-1], result[i][j-1])
			result[i][j] = t + 1
			if result[i][j] > maxSize {
				maxSize = result[i][j]
			}
		}
	}

	return maxSize
}

func min(a, b, c int) int {
	l := a
	if b < l {
		l = b
	}
	if c < l {
		l = c
	}
	return l
}

func main() {
	arr := [][]int{{0, 1, 1, 0, 1}, {1, 1, 1, 0, 0}, {1, 1, 1, 1, 0}, {1, 1, 1, 0, 1}}
	result := MaximumSizeSubMatrix(arr)
	fmt.Println(result)
}
