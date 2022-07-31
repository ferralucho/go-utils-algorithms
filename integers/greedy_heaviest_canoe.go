package main

import "fmt"

/*
Problem: There are n > 0 canoeists weighing respectively 1 <= 10 TO THE 9
The goal is to seat them in the minimum number of double canoes whose displacement (the
maximum load) equals k. You may assume that w i <= k.
*/

func greedyCanoes(weights []int, k int) int {
	canoes := 0
	j := 0
	i := len(weights) - 1
	for i >= j {
		if weights[i]+weights[j] <= k {
			j += 1
		}
		canoes += 1
		i -= 1
	}
	return canoes
}

func main() {
	weights := []int{25, 35, 43, 75, 99, 79, 41}
	loadFactor := 70
	fmt.Println(greedyCanoes(weights, loadFactor))
}
