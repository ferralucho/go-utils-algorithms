/*
 * Problem: Merge Intervals
 *
 * Given a collection of intervals, merge all overlapping intervals.
 *
 * Example 1:
 *
 * Input: [[1,3],[15,18],[8,10],[2,6]]
 * Output: [[1,6],[8,10],[15,18]]
 *
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
 *
 * Example 2:
 *
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 *
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 */

// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(MergeIntervals(intervals))
}

func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return i < j
	})

	var newIntervals [][]int
	for i := 0; i < len(intervals); i++ {
		lower := intervals[i][0]
		upper := intervals[i][1]
		if len(intervals)-i-1 > 0 {
			next := intervals[i+1]
			if upper >= next[0] {
				element := []int{
					lower, next[1],
				}

				newIntervals = append(newIntervals, element)
				i += 1
			} else {
				newIntervals = addElement(lower, upper, newIntervals)
			}
		} else {
			newIntervals = addElement(lower, upper, newIntervals)
		}
	}

	return newIntervals
}

func addElement(lower int, upper int, newIntervals [][]int) [][]int {
	element := []int{
		lower, upper,
	}
	newIntervals = append(newIntervals, element)
	return newIntervals
}
