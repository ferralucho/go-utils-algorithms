/*
Use Dynamic Programming for Optimization Problems
- DP breaks problems into smaller overlapping sub-problems for optimization.
- It's often used for maximization, minimization, or counting paths.
- Example: Solve the 0/1 knapsack problem.
*/

package main

import "fmt"

func knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7
	maxValue := knapsack(weights, values, capacity)
	fmt.Println("Maximum value:", maxValue) // Output: Maximum value: 9
}
