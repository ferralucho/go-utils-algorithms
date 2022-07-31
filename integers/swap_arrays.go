package main

/*
Problem: You are given an integer m (1 < m <= 1 000 000) and two non-empty, zero-indexed
arrays A and B of n integers, a 0 , a 1 , . . . , a n−1 and b 0 , b 1 , . . . , b n−1 respectively (0 � a i , b i � m).
The goal is to check whether there is a swap operation which can be performed on these
arrays in such a way that the sum of elements in array A equals the sum of elements in
array B after the swap. By swap operation we mean picking one element from array A and
one element from array B and exchanging them.
Solution O(n 2 ): The simplest method is to swap every pair of elements and calculate the
totals. Using that approach gives us O(n 3 ) time complexity. A better approach is to calculate
the sums of elements at the beginning, and check only how the totals change during the swap
operation.
*/

func SlowSwapArraysSum(a, b []int) bool {
	n := len(a)
	sumA := sum(a)
	sumB := sum(b)
	var change int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			change = b[j] - a[i]
			sumA += change
			sumB -= change
			if sumA == sumB {
				return true
			}
			sumA -= change
			sumB += change
		}
	}
	return false

}

func sum(toSum []int) int {
	sum := 0
	for i := 0; i < len(toSum); i++ {
		sum += toSum[i]
	}
	return sum
}

/*
func main() {
	var a = []int{1, 8}
	var b = []int{2, 7}

	println(SlowSwapArraysSum(a, b) == false)

	var c = []int{1, 5}
	var d = []int{1, 5}
	println(SlowSwapArraysSum(c, d) == true)
}
*/
