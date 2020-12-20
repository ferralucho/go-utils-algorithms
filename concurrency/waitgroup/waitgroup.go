package waitgroup

import (
	"fmt"
	"sync"
)

func doSomething(wg *sync.WaitGroup) {
	// do something here
	fmt.Println("Done")
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go doSomething(&wg)

	// program will wait until doSomething & doSomethingElseSync is complete
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}
		}
	}()
	done <- true
	fmt.Println("Done channel")
}
