/*
3. Backtracking or BFS for Combinations
- Use Backtracking or BFS to explore all combinations or permutations.
- They’re great for generating subsets or solving puzzles.
- Example: Generate all possible subsets of a given set.
*/

package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{{}} // Start with an empty subset

	for _, num := range nums {
		n := len(result)
		for i := 0; i < n; i++ {
			newSubset := make([]int, len(result[i]))
			copy(newSubset, result[i])
			newSubset = append(newSubset, num)
			result = append(result, newSubset)
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 7, 8}
	subsets := subsets(nums)
	fmt.Println(subsets)
}
