package main

import "fmt"

func main() {
	arr := []int{5, 9, 2, 6, 1, 8, 3, 7}
	fmt.Println("Original array:", arr)
	MergeSort(arr)
	fmt.Println("Sorted array:", arr)
}

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	MergeSort(left)
	MergeSort(right)
	Merge(arr, left, right)
}

func Merge(arr []int, left []int, right []int) {
	i := 0 // index for left subarray
	j := 0 // index for right subarray
	k := 0 // index for merged array

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
