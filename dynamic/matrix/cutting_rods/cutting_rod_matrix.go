package main

import (
	"fmt"
	"math"
)

/**
 * http://www.geeksforgeeks.org/dynamic-programming-set-13-cutting-a-rod/
 */

func MaxValue(price []int) int {
	max := make([]int, len(price)+1)
	for i := 1; i <= len(price); i++ {
		for j := i; j <= len(price); j++ {
			max[j] = int(math.Max(float64(max[j]), float64(max[j-i]+price[i-1])))
			fmt.Println("max[j]", max[j], "max[j-i]", max[j-i], "price[i-1]", price[i-1])
		}
	}
	return max[len(price)]
}

func main() {
	prices := []int{3, 5, 8, 9, 10, 20, 22, 25}
	fmt.Println("MaxValue", MaxValue(prices))
}
