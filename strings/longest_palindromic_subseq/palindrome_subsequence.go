package main

import (
	"fmt"
)

/*
 *
 * Given a string find longest palindromic subsequence in this string.
 *
 * Time complexity - O(n2)
 * Space complexity - O(n2
 *
 *
 * References
 * http://www.geeksforgeeks.org/dynamic-programming-set-12-longest-palindromic-subsequence/
 */

func CalculateRecursive(str []byte, start, length int) int {
	if length == 1 {
		return 1
	}
	if length == 0 {
		return 0
	}
	if str[start] == str[start+length-1] {
		return 2 + CalculateRecursive(str, start+1, length-2)
	} else {
		return max(CalculateRecursive(str, start+1, length-1), CalculateRecursive(str, start, length-1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str := []byte{'a', 'b', 'c', 'a', 'd', 'e', 'c', 'f', 'a'}
	result := CalculateRecursive(str, 0, len(str))
	fmt.Println(result)
}
