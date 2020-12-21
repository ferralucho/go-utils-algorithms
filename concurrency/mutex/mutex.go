package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 0
	var wg sync.WaitGroup

	var mu sync.Mutex // guards access

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			a += 1
		}()
	}
	wg.Wait()
	fmt.Println(a) // will always be 1000
}
