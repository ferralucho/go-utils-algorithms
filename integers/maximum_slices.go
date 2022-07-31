package main

import "fmt"

/*
Let’s deﬁne a problem relating to maximum slices. You are given a sequence of n integers
a 0 , a 1 , . . . , a n−1 and the task is to ﬁnd the slice with the largest sum. More precisely, we are
looking for two indices p, q such that the total a p + a p+1 + . . . + a q is maximal. We assume
that the slice can be empty and its sum equals 0.
a 0 a 1 a 2 a 3 a 4 a 5 a 6
5 -7 3 5 -2 4 -1
The sum of this slice
equals 10 and there is no slice with a larger sum. Notice that the slice we are looking for may
contain negative integers, as shown above.
*/

func MaximumSlices(nums []int) int {
	maxEnding, maxSlice := 0, 0
	for _, v := range nums {
		maxEnding = maxSum(0, maxEnding+v)
		maxSlice = maxSum(maxSlice, maxEnding)
	}
	return maxSlice
}

func maxSum(first, second int) int {
	if first >= second {
		return first
	}
	return second
}

func main() {
	nums := []int{5, -7, 3, 5, -2, 4, -1}
	fmt.Println(MaximumSlices(nums))
}
