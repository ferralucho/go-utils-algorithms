package main

import "fmt"

// LongestCommonSubstring calculates the length of the longest common substring
// between two strings using dynamic programming.
func LongestCommonSubstring(str1, str2 []rune) int {
	m := len(str1)
	n := len(str2)
	T := make([][]int, m+1)
	for i := range T {
		T[i] = make([]int, n+1)
	}

	maxLength := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				T[i][j] = T[i-1][j-1] + 1
				if maxLength < T[i][j] {
					maxLength = T[i][j]
				}
			}
		}
	}

	return maxLength
}

// LongestCommonSubstringRec calculates the length of the longest common substring
// between two strings using recursive approach.
func LongestCommonSubstringRec(str1, str2 []rune, pos1, pos2 int, checkEqual bool) int {
	if pos1 == -1 || pos2 == -1 {
		return 0
	}
	if checkEqual {
		if str1[pos1] == str2[pos2] {
			return 1 + LongestCommonSubstringRec(str1, str2, pos1-1, pos2-1, true)
		}
		return 0
	}

	r1 := 0
	if str1[pos1] == str2[pos2] {
		r1 = 1 + LongestCommonSubstringRec(str1, str2, pos1-1, pos2-1, true)
	}

	return max(r1, max(LongestCommonSubstringRec(str1, str2, pos1-1, pos2, false), LongestCommonSubstringRec(str1, str2, pos1, pos2-1, false)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str1 := []rune("abcdef")
	str2 := []rune("zcdemf")

	fmt.Println(LongestCommonSubstring(str1, str2))
	fmt.Println(LongestCommonSubstringRec(str1, str2, len(str1)-1, len(str2)-1, false))
}
