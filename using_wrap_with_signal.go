package main

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func main() {
    signal:= make(chan bool)
    go WrapWithSignal(play, signal)()

    if <-signal{
        fmt.Println("Congratulations, you win!")
        return
    }

    fmt.Println("You died.")
}

func play(signal chan bool){
    fmt.Println("Playing russian roulette")
    rand.Seed(time.Now().UnixNano())

    if rand.Intn(2) == 1{
        panic(errors.New("got a bullet"))
    }

    signal <- true
}
