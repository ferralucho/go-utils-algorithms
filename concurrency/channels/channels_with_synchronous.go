package main

import (
	"fmt"
)

func main() {
	var numbers []int // nil
	done := make(chan struct{})
	// start a goroutine to initialise array
	go func() {
		numbers = make([]int, 2)
		done <- struct{}{}
	}()

	// do something synchronous
	<-done                  // read done from channel
	numbers[0] = 1          // will not panic anymore
	fmt.Println(numbers[0]) // 1
}
