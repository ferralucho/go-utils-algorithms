package accounts

import (
	"errors"
	"my-package/amounts"
)

type Account struct {
	balance amounts.Amount
}

func NewEmptyAccount() Account {
	amount, _ := amounts.NewAmount(0)
	return NewAccount(amount)
}

func NewAccount(amount amounts.Amount) Account {
	return Account{balance: amount}
}

func (acc Account) Deposit(amount amounts.Amount) Account {
	newAmount, _ := amounts.NewAmount(acc.balance.GetValue() + amount.GetValue())
	acc.balance = newAmount
	return acc
}

func (acc Account) Withdraw(amount amounts.Amount) (Account, error) {
	newAmount, err := amounts.NewAmount(acc.balance.GetValue() - amount.GetValue())
	if err != nil {
		return acc, errors.New("Insuficient funds")
	}
	acc.balance = newAmount
	return acc, nil
}
