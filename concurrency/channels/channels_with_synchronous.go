package main

import (
	"fmt"
)

//Channels can be buffered if you want to prevent blocking further execution until a value is
//eventually read from the channel to free it up.

//Channels can be non-buffered if you want 1-in-1-out behavior.

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
	fmt.Println(numbers[1]) // 0
}
