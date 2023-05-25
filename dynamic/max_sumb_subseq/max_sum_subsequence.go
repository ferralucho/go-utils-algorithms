package main

/**
 * Problem Statement:
 *
 * Given an array of n positive integers. Write a program to find the sum of maximum sum subsequence of the given array
 * such that the integers in the subsequence are in increasing order.
 *
 *
 * Video: https://youtu.be/99ssGWhLPUE
 *
 * Reference:
 * http://www.geeksforgeeks.org/dynamic-programming-set-14-maximum-sum-increasing-subsequence/
 */

import "fmt"

func MaxSum(arr []int) int {
	T := make([]int, len(arr))

	for i := 0; i < len(T); i++ {
		T[i] = arr[i]
	}

	for i := 1; i < len(T); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				T[i] = max(T[i], T[j]+arr[i])
			}
		}
	}

	max := T[0]
	for i := 1; i < len(T); i++ {
		if T[i] > max {
			max = T[i]
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 101, 10, 2, 3, 100, 4}
	result := MaxSum(arr)
	fmt.Println(result)
}
