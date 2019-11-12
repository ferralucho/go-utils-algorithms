package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

//Panic with custom error
func ThrowCustomError(name string, msg string) {
	panic(&customError{name: name, err: msg})
}

//Create an error interface, we handle our own errors only
type CustomError interface {
	error
	Name() string
}

//Custom error that will be handled
type customError struct {
	name string
	err  string
}

func (e *customError) Error() string {
	return e.err
}
func (e *customError) Name() string {
	return e.name
}

//Panic group that will store the panic errors
type PanicGroup struct {
	wg      sync.WaitGroup
	errOnce sync.Once
	err     []CustomError
}

func (g *PanicGroup) Wait() []CustomError {
	g.wg.Wait()
	return g.err
}
func (g *PanicGroup) Go(f func()) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(CustomError); ok {
					g.err = append(g.err, err)
				} else {
					log.Fatal("[FATAL]", r)
				}
			}
		}()
		f()
	}()
}

//Generate a negative or positive float for random test data
func TestData() (float64, float64) {
	rand.Seed(time.Now().UnixNano())
	var f1, f2 = -23.44, 5.44
	if rand.Intn(2) == 0 {
		f1 = f2
	}
	if rand.Intn(2) == 0 {
		f2 = f1
	}
	return f1, f2
}

//Get the Sqlt of a float
func Sqrt(f float64) float64 {
	if f < 0 {
		ThrowCustomError("calc_sqrt_negative", "Can't calc sqrt of a negative number")
	}
	return math.Sqrt(f)
}

func main() {
	var g PanicGroup
	f1, f2 := TestData()

	// Goroutine #1
	g.Go(func() {
		fmt.Printf("[result] %f\n", Sqrt(f1))
	})

	// Goroutine #2
	g.Go(func() {
		fmt.Printf("[result] %f\n", Sqrt(f2))
	})

	// Wait till all goroutines are done and handle errors
	for _, err := range g.Wait() {
		HandleErr(err)
	}
}

//Handle error
func HandleErr(err CustomError) {
	log.Printf("[%s] %s", err.Name(), err.Error())
}
