package amounts

import "errors"

type Amount struct {
	value int
}

func NewAmount(value int) (Amount, error) {
	if value < 0 {
		return Amount{}, errors.New("Invalid amount")
	}

	return Amount{value: value}, nil
}

func (a Amount) GetValue() int {
	return a.value
}
