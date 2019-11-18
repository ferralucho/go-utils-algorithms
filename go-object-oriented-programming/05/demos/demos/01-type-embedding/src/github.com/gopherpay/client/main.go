package main

import "fmt"

type Account struct{}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds")
	return 0
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

type CreditAccount struct {
	Account
}

func main() {
	ca := &CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)

}
