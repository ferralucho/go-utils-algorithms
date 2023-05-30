package main

import "fmt"

func binarySearch(arr []int, low, high, key int) bool {
	// Base Case
	if low > high {
		return false
	}

	// Calculating Mid
	mid := (low + high) / 2

	// Checking if the key is found in the middle
	if arr[mid] == key {
		return true
	}

	// Searching on the left half if a key exists there
	if key < arr[mid] {
		return binarySearch(arr, low, mid-1, key)
	}

	// Searching on the other half otherwise
	return binarySearch(arr, mid+1, high, key)
}

func main() {
	arr := []int{2, 5, 9, 13, 17, 21, 30}
	if binarySearch(arr, 0, len(arr)-1, 30) {
		fmt.Println("Element Found")
	} else {
		fmt.Println("Element not Found")
	}
}
