package main

import "fmt"

/*
http://www.geeksforgeeks.org/program-nth-catalan-number/
Count number of binary search tree created for array of size n
Use of catalan number
*/

func countTreesRecursive(numKeys int) int {
	if numKeys <= 1 {
		return 1
	} else {
		sum := 0
		var left, right, root int
		for root = 1; root <= numKeys; root++ {
			left = countTreesRecursive(root - 1)
			right = countTreesRecursive(numKeys - root)
			sum += left * right
		}
		return sum
	}
}

func countTrees(n int) int {
	T := make([]int, n+1)
	T[0] = 1
	T[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			T[i] += T[j] * T[i-j-1]
		}
	}
	return T[n]
}

func main() {
	numberOfTrees := countTrees(5)
	fmt.Println(numberOfTrees)

	recTrees := countTreesRecursive(5)
	fmt.Println("recursive trees", recTrees)
}
