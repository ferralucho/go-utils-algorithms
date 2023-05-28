package main

/*
Problem:
- Given an array of integers, return indices of the two numbers such that they
  add up to a specific target.
- You may assume that each input would have exactly one solution, and you may
  not use the same element twice.

Example:
- Input: nums = []int{2, 5, 4}, target = 6
  Output: [0, 2] since nums[0] + nums[2] = 2 + 4 = 6

Approach:
- Use a hash map to store the value and its index as we iterate through the
  list.
- Within each iteration, look up the difference of target and the current
  value to see if we have seen that number.
- Simply return two cached indices once that condition meets.

Cost:
- O(n) time, O(n) space.
*/

func FindSum(nums []int, target int) []int {
	m := map[int]int{}
	out := make([]int, 2)

	// return empty slice on edge cases.
	if len(nums) == 0 || len(nums) == 1 {
		return out
	}

	// iterate through the list, look up the temporary hash map to see if we
	// have seen the number before and return both cached indices.
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			out[0] = j
			out[1] = i
		}

		m[num] = i
	}

	return out
}

/*
func main() {
	fmt.Println("test")
}
*/
