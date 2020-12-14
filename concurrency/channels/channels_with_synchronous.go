package main

import (
	"fmt"
)

/*
Using a channel, we can make sure the main task waits until the asynchronous task is complete.
When the goroutine completes its’ work, it will send a value through the channel done , which will be read before operating on the numbers array.
Channels can be buffered if you want to prevent blocking further execution until a value is
eventually read from the channel to free it up.

Channels can be non-buffered if you want 1-in-1-out behavior.
*/

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

/*
package main

import (
	"fmt"
	"sync"
)

//Data race
func main() {
	a := 0 // data race
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			a += 1
		}()
	}
	wg.Wait()
  	fmt.Println(a) // could theoretical be any number 0-1000 (most likely above 900)
}
*/
