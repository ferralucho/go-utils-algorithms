package main

import "fmt"

/*
	Given a total and coins of certain denominations find number of ways total
	can be formed from coins assuming infinity supply of coins

	http://www.geeksforgeeks.org/dynamic-programming-set-7-coin-change/
*/

func PrintCoinChangingSolution(total int, coins []int) {
	result := make([]int, 0)
	printActualSolution(&result, total, coins, 0)
}

func printActualSolution(result *[]int, total int, coins []int, pos int) {
	if total == 0 {
		for _, r := range *result {
			fmt.Printf("%d ", r)
		}
		fmt.Println()
	}
	for i := pos; i < len(coins); i++ {
		if total >= coins[i] {
			*result = append(*result, coins[i])
			printActualSolution(result, total-coins[i], coins, i)
			*result = (*result)[:len(*result)-1]
		}
	}
}

func main() {
	total := 15
	coins := []int{1, 2, 5, 10}
	PrintCoinChangingSolution(total, coins)
}
