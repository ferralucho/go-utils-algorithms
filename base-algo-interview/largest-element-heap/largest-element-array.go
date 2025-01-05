package main

/*
Use a Heap for K Elements
- When finding the top K largest or smallest elements, heaps are your best tool.
- They efficiently handle priority-based problems with O(log K) operations.
- Example: Find the 3 largest numbers in an array.

*/
/* import (
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
*/

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKLargest(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)

	// Push the first 'k' elements into the heap
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	// Iterate through the remaining elements
	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0] { // If current number is greater than the smallest in the heap
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	// Extract the 'k' largest elements from the heap
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(int)
	}
	return result
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 3
	result := findKLargest(nums, k)
	fmt.Println("Top", k, "largest numbers:", result) // Output: Top 3 largest numbers: [6 5 4]
}
