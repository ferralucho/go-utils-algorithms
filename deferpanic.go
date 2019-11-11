package main

import 
(
"fmt"
"time"
)

func main() {
	go Wrap(test)()
 	time.Sleep(time.Second)
 	fmt.Println("HELLO")
}

func test(){
 	panic("PANIC")
}

func Wrap(f func()) func(){return func(){
 	defer func(){if r := recover(); r !=nil{
 	fmt.Printf("RECOVERED - %v\r\n", r)}}()
 	f()}
}

func WrapWithSignal(f func(chan bool), signal chan bool) func(){return func(){
	 defer func(){if r := recover(); r !=nil{
	 fmt.Printf("RECOVERED - %v\r\n", r)
	 signal <-false}}()
	 f(signal)}
 }
