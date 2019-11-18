package objectoriented

type PaymentOption interface {
	ProcessPayment(float32) bool
}
