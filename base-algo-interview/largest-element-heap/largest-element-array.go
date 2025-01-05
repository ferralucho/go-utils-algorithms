package main

import (
	"fmt"
	"sort"
)

func findKLargest(n []int, k int) []int {
	sort.Slice(n, func(i, j int) bool {
		return n[i] > n[j]
	})

	return n[:k]
}

func main() {
	n := []int{3, 1, 2, 6, 4, 5}
	k := 2
	result := findKLargest(n, 2)
	fmt.Println("Top", k, "largest numbers:", result)
}
