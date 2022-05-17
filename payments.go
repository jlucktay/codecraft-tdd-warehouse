//go:generate mockgen -source=payments.go -destination=./mock_payments_test.go -package=warehouse_test

package warehouse

type PaymentProvider interface {
	Sell(price float64) error
	Refund(price float64) error
}
