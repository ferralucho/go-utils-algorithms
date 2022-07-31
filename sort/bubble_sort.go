package main

import "fmt"

func main() {
	var numbers []int = []int{78, 95, 65, 99, 24}
	fmt.Println("Unsorted:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		sweep(numbers, i)
	}
}

func sweep(numbers []int, prevPasses int) {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	fmt.Println("Steps:", numbers)

	for secondIndex < (N - prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}
