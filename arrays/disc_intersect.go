package main

import (
	"fmt"
	"sort"
)

func GetNumberDiscInsersections(A []int) int {
	leftDiscs := make([]int, len(A))
	rightDiscs := make([]int, len(A))
	result := 0

	for i := 0; i < len(A); i += 1 {
		leftDiscs[i] = i - A[i]
		rightDiscs[i] = i + A[i]
	}

	sort.Ints(leftDiscs)
	sort.Ints(rightDiscs)

	fmt.Println("left discs", leftDiscs)
	fmt.Println("right discs", rightDiscs)
	for i := 0; i < len(A); i += 1 {
		end := rightDiscs[i]

		count := sort.Search(len(leftDiscs), func(i int) bool {
			return leftDiscs[i] > end
		})

		count = count - (i + 1)
		result += count

		upperLimitChecked, done := checkUpperLimit(result)
		if done {
			return upperLimitChecked
		}
	}
	return result
}

func checkUpperLimit(result int) (int, bool) {
	if result > 10000000 {
		return -1, true
	}
	return 0, false
}

func main() {
	discPoints := []int{1, 5, 2, 1, 4, 0}
	fmt.Println(GetNumberDiscInsersections(discPoints))
}
