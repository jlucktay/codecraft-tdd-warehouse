package warehouse

type paymentProvider interface {
	Sell(price float64) error
	Refund(price float64) error
}

type dummyPayments struct{}

func (dp *dummyPayments) Sell(price float64) error {
	return nil
}

func (dp *dummyPayments) Refund(price float64) error {
	panic("not implemented") // TODO: Implement
}
