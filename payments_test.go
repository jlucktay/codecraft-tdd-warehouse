package warehouse_test

type dummyPayments struct{}

func (dp *dummyPayments) Sell(price float64) error {
	return nil
}

func (dp *dummyPayments) Refund(price float64) error {
	panic("not implemented") // TODO: Implement
}
