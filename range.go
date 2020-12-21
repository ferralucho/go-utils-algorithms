package main

import "fmt"

var pow = []int{2, 4, 8, 16}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
