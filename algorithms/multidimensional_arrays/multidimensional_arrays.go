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
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		currStart := intervals[i][0]
		currEnd := intervals[i][1]

		if currStart <= end {
			if currEnd > end {
				end = currEnd
			}
		} else {
			merged = append(merged, []int{start, end})
			start = currStart
			end = currEnd
		}
	}

	merged = append(merged, []int{start, end})

	return merged
}
