/*
You are given an array of non-negative integers numbers. You are allowed to choose any number from this array and swap any two digits in it. If after the swap operation the number contains leading zeros, they can be omitted and not considered (eg: 010 will be considered just 10).

Your task is to check whether it is possible to apply the swap operation at most once, so that the elements of the resulting array are strictly increasing.

# Example

For numbers = [1, 5, 10, 20], the output should be solution(numbers) = true.

The initial array is already strictly increasing, so no actions are required.

For numbers = [1, 3, 900, 10], the output should be solution(numbers) = true.

By choosing numbers[2] = 900 and swapping its first and third digits, the resulting number 009 is considered to be just 9. So the updated array will look like [1, 3, 9, 10], which is strictly increasing.

For numbers = [13, 31, 30], the output should be solution(numbers) = false.

The initial array elements are not increasing.
By swapping the digits of numbers[0] = 13, the array becomes [31, 31, 30] which is not strictly increasing;
By swapping the digits of numbers[1] = 31, the array becomes [13, 13, 30] which is not strictly increasing;
By swapping the digits of numbers[2] = 30, the array becomes [13, 31, 3] which is not strictly increasing;
So, it's not possible to obtain a strictly increasing array, and the answer is false.

Input/Output

[execution time limit] 4 seconds (go)

[memory limit] 1 GB

[input] array.integer numbers

An array of non-negative integers.

Guaranteed constraints:
1 ≤ numbers.length ≤ 103,
0 ≤ numbers[i] ≤ 103.

[output] boolean

Return true if it is possible to obtain a strictly increasing array by applying the digit-swap operation at most once, and false otherwise.
*/

package main

import (
	"fmt"
	"strconv"
)

func check(xs []int) bool {
	x := strconv.Itoa(xs[1])
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			numStr := x[:i] + string(x[j]) + x[i+1:j] + string(x[i]) + x[j+1:]
			num, _ := strconv.Atoi(numStr)
			if xs[0] < num && num < xs[2] {
				return true
			}
		}
	}
	return false
}

func strictInc(xs []int, singleFlipAllowed bool) bool {
	for n := 0; n < len(xs)-1; n++ {
		if xs[n] >= xs[n+1] {
			if singleFlipAllowed {
				if n > 0 {
					if check([]int{xs[n-1], xs[n], xs[n+1]}) && strictInc(xs[n+1:], false) {
						return true
					}
				}
				if n < len(xs)-2 {
					if check([]int{xs[n], xs[n+1], xs[n+2]}) && strictInc(xs[n+2:], false) {
						return true
					}
				}
			}
			return false
		}
	}
	return true
}

func main() {
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println("1, 2, 3, 4, 5", strictInc(xs, true)) // true

	xs = []int{5, 4, 3, 2, 1}
	fmt.Println("5, 4, 3, 2, 1", strictInc(xs, true)) // false

	xs = []int{123, 321, 456, 789}
	fmt.Println("123, 321, 456, 789", strictInc(xs, true)) // true

	xs = []int{123, 321, 789, 456}
	fmt.Println("123, 321, 789, 456", strictInc(xs, true)) // false

	xs = []int{12, 21, 34, 45}
	fmt.Println("12, 21, 34, 45", strictInc(xs, true)) // true

	xs = []int{12, 21, 43, 34}
	fmt.Println("12, 21, 43, 34", strictInc(xs, true)) // false
}
