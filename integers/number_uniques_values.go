package main

import (
	"fmt"
	"sort"
)

func NumberUniqueValues(nums []int) int {
	sort.Ints(nums)
	result := 1
	for i, v := range nums {
		fmt.Println(i)

		if i > 0 && v != nums[i-1] {
			result += 1
		}
	}
	return result
}

func main() {
	var a = []int{1, 83, 7, 11, 2, 3, 3}
	println(NumberUniqueValues(a))
}
