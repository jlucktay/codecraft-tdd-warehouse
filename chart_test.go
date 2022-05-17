package warehouse_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matryer/is"
	warehouse "go.jlucktay.dev/codecraft-tdd-warehouse"
)

type StubChart struct{}

func (sc *StubChart) Notify(artist string, title string, copies uint) {
	fmt.Println("stub chart notified")
}

type StubPriceGuarantor struct{}

func (spg *StubPriceGuarantor) BestPrice(artist, title string) (price float64, err error) {
	return 42.0, nil
}

func TestChartIsNotifiedOfSale(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	m := NewMockChart(ctrl)

	m.
		EXPECT().
		Notify(gomock.Any(), gomock.Any(), gomock.Any()).
		Return()

	whse := warehouse.New(m, &dummyPayments{})
	newCD := warehouse.NewCD("Metallica", "title string", "label string", 42.0, 27)

	newID, err := whse.Add(newCD)
	is.NoErr(err) // Could not add new CD to warehouse

	is.NoErr(whse.SellSingleCD(newID)) // Could not sell single CD
}
