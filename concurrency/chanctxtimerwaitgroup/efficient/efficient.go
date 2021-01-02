package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
Context’s Done Channel

Using ctx.Done() actually performs mutex locks and unlocks behind the scenes to retrieve the done channel.
We can retrieve the done channel prior to using it once and avoid all the mutex locks and unlocks.

We can reuse the same timer instead of creating a new one every time it fires,
we just need to remember to reset the timer when it fires.
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Caught %s\n", sig)
		cancel()
	}()

	filename := "test.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Can't load %s - %v\n", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make(chan string, 1)

	done := ctx.Done()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		defer close(lines)
		for {
			if !scanner.Scan() {
				fmt.Println("Reader: Completed")
				break
			}
			lines <- scanner.Text()
			time.Sleep(time.Second)

			select {
			case <-done:
				fmt.Println("Reader: Context closed")
				return ctx.Err()
			default:
			}
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("reader error - %w", err)
		}
		return nil
	})

	timeout := time.Millisecond * 500
	timeoutTimer := time.NewTimer(timeout)
	defer timeoutTimer.Stop()
	eg.Go(func() error {
		for {
			select {
			case <-done:
				fmt.Println("Sender: Context closed")
				return ctx.Err()
			case <-timeoutTimer.C:
				timeoutTimer.Reset(timeout)
				fmt.Println("Sender: Write to channel is taking sometime")
			case l, ok := <-lines:
				if !ok {
					fmt.Printf("Sender: Channel closed\n", l)
					return nil
				}
				fmt.Printf("Sender: Sending %s to remote database\n", l)
			}
		}
	})

	err = eg.Wait()
	if err != nil {
		fmt.Printf("error from goroutine - %v\n", err)
	}
}
