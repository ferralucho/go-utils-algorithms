package objectoriented

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (c Cash) ProcessPayment(amount float32) bool {
	return true
}
