package main

import (
	"fmt"
	"math"
)

/**
 * Given an array of non negative integers where each element represents the max
 * number of steps that can be made forward from that element. Write a function to
 * return the minimum number of jumps to reach the end of the array from first element

 * Solution
 * Have 2 for loop. j trails i. If arr[j] + j >= i then you calculate new jump
 * and result.

 * Space complexity O(n) to maintain result and min jumps
 * Time complexity O(n^2)

 * Reference
 * http://www.geeksforgeeks.org/minimum-number-of-jumps-to-reach-end-of-a-given-array/
 */

func minimumJump(arr []int, results []int) int {
	jumps := make([]int, len(arr))
	jumps[0] = 0

	for i := 1; i < len(arr); i++ {
		jumps[i] = math.MaxInt - 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j]+j >= i {
				if jumps[i] > jumps[j]+1 {
					results[i] = j
					jumps[i] = jumps[j] + 1
				}
			}
		}
	}

	return jumps[len(jumps)-1]
}

func main() {
	arr := []int{1, 3, 5, 3, 2, 2, 6, 1, 6, 8, 9}
	results := make([]int, len(arr))
	min := minimumJump(arr, results)
	fmt.Println(min)
}
