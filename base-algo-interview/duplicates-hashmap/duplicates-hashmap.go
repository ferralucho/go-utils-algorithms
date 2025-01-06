/*
Optimize Arrays with HashMaps or Sorting
- Replace nested loops with HashMaps for O(n) solutions or sorting for O(n log n).
- HashMaps are perfect for lookups, while sorting simplifies comparisons.
- Example: Find duplicates in an array.
*/

package main

import "fmt"

func findDuplicates(nums []int) []int {
	result := []int{}
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	duplicates := findDuplicates(nums)
	fmt.Println("Duplicates:", duplicates) // Output: Duplicates: [2 3]
}
