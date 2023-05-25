package main

import (
	"fmt"
	"math"
)

/**
 * Find a subsequence in given array in which the subsequence's elements are
 * in sorted order, lowest to highest, and in which the subsequence is as long as possible
 *
 * Solution :
 * Dynamic Programming is used to solve this question. DP equation is
 * if(arr[i] > arr[j]) { T[i] = max(T[i], T[j] + 1 }
 *
 * Time complexity is O(n^2).
 * Space complexity is O(n)
 *
 * Reference
 * http://en.wikipedia.org/wiki/Longest_increasing_subsequence
 * http://www.geeksforgeeks.org/dynamic-programming-set-3-longest-increasing-subsequence/
 */

func LongestSubSequence(arr []int) int {
	T := make([]int, len(arr))
	actualSolution := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		T[i] = 1
		actualSolution[i] = i
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				if T[j]+1 > T[i] {
					T[i] = T[j] + 1
					actualSolution[i] = j
				}
			}
		}
	}

	maxIndex := 0
	for i := 0; i < len(T); i++ {
		if T[i] > T[maxIndex] {
			maxIndex = i
		}
	}

	t := maxIndex
	newT := maxIndex
	for t != newT {
		t = newT
		fmt.Printf("%d ", arr[t])
		newT = actualSolution[t]
	}
	fmt.Println()

	return T[maxIndex]

}

func LongestSubSeqRecursive(arr []int) int {
	maxLen := 0
	for i := 0; i < len(arr); i++ {
		longl := longestSubSeqRecursive(arr, i+1, arr[i])
		if longl > maxLen {
			maxLen = longl
		}
	}
	return maxLen + 1
}

func longestSubSeqRecursive(arr []int, pos, lastNum int) int {
	if pos == len(arr) {
		return 0
	}
	t1 := 0
	if arr[pos] > lastNum {
		t1 = 1 + longestSubSeqRecursive(arr, pos+1, arr[pos])
	}
	t2 := longestSubSeqRecursive(arr, pos+1, lastNum)
	return int(math.Max(float64(t1), float64(t2)))
}

func main() {
	arr := []int{23, 10, 22, 5, 33, 8, 9, 21, 50, 41, 60, 80, 99, 22, 23, 24, 25, 26, 27}
	resRec := LongestSubSeqRecursive(arr)
	fmt.Println("result:", resRec)

	resDyn := LongestSubSequence(arr)
	fmt.Println("dynamic result:", resDyn)
}
