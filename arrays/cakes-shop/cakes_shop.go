package main

import (
	"fmt"
	"strconv"
)

func MinDaysWaitCake(n int, m int, k int, day []int) int {
	//n available suppliers
	//m cakes
	//k adjacent suppliers

	maxDay := 1

	for _, k := range day {
		maxDay = max(maxDay, k)
	}

	for i := 1; i <= maxDay; i++ {
		arr := make([]int, 0)
		totalCake := m
		//iterate suppliers
		for j := 0; j < n; j++ {
			if day[j] <= i {
				arr = append(arr, 0)
			} else {
				arr = append(arr, 1)
			}
		}

		count := 0
		for j := 0; j < n; j++ {
			if arr[j] == 1 {
				count++
			} else {
				totalCake -= (count / k)
				count = 0
			}
		}
		totalCake = (count / k)

		if totalCake <= 0 {
			return i
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	const suppliers = 7
	cakes := 2
	adjacentSuppliers := 3
	days := []int{7, 7, 12, 7, 12, 7, 12}

	fmt.Printf(strconv.Itoa(MinDaysWaitCake(suppliers, cakes, adjacentSuppliers, days)))
}
