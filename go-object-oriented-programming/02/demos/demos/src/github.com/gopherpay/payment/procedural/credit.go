package procedural

type CreditCard struct {
	OwnerName       string
	CardNumber      string
	ExpirationMonth int
	ExpirationYear  int
	SecurityCode    int
	AvailableCredit float32
}

func PayWithCredit(card *CreditCard, amount float32) bool {
	if card.AvailableCredit > amount {
		card.AvailableCredit -= amount
		return true
	} else {
		return false
	}
}
