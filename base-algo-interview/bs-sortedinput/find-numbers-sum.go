/* 2. Binary Search or Two Pointers for Sorted Inputs
- Sorted arrays often point to Binary Search or Two Pointer techniques.
- These methods drastically reduce time complexity to O(log n) or O(n).
- Example: Find two numbers in a sorted array that add up to a target.
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{} // No solution found
}

func main() {
	nums := []int{-2, -1, 1, 1, 2, 3}
	target := 0
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [1 4]
}
