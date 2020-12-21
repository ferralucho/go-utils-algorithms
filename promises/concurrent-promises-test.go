package main
 
import (
   "testing"
   "time"
   "fmt"
)
 
func TestThePromisesRunConcurrently(t *testing.T) {
   sleepOneSecond := func() (interface{}, error){
      fmt.Println("Se ejecuta!!")
      time.Sleep(1000 * time.Millisecond)
      return nil, nil
   }
 
   start := time.Now()
 
   promise1 := New(sleepOneSecond)
   promise2 := New(sleepOneSecond)
   promise3 := New(sleepOneSecond)
   promise4 := New(sleepOneSecond)
   promise5 := New(sleepOneSecond)
 
   promise1.Get()
   promise2.Get()
   promise3.Get()
   promise4.Get()
   promise5.Get()
 
   end := time.Now()
 
   timeElapsed := end.Sub(start)
 
   if timeElapsed.Seconds() <= 1 || timeElapsed.Seconds() >= 2 {
      t.Error("It should have done it in 1 second aprox")
   }
}
