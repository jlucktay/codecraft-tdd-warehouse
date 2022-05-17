package warehouse

type PriceGuarantor interface {
	BestPrice(artist, title string) (price float64, err error)
}
