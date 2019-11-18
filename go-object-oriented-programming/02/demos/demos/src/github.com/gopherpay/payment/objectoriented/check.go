package objectoriented

type CheckingAccount struct {
	accountOwner  string
	routingNumber string
	accountNumber string
	balance       float32
}

func CreateCheckingAccount(accountOwner, routingNumber, accountNumber string) *CheckingAccount {
	return &CheckingAccount{
		accountOwner:  accountOwner,
		routingNumber: routingNumber,
		accountNumber: accountNumber,
		balance:       250, // this should come from a call to a web service
	}
}

func (c CheckingAccount) ProcessPayment(amount float32) bool {
	return true
}

func (c CheckingAccount) Balance() float32 {
	return c.balance
}
