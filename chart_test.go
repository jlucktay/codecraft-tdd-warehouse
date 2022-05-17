package warehouse_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matryer/is"
	warehouse "go.jlucktay.dev/codecraft-tdd-warehouse"
)

func TestChartIsNotifiedOfSale(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	m := NewMockChart(ctrl)

	m.
		EXPECT().
		Notify(gomock.Any(), gomock.Any(), gomock.Any()).
		Return()

	whse := warehouse.New(m)
	newCD := warehouse.NewCD("title string", "label string", 42.0, 27)

	newID, err := whse.Add(newCD)
	is.NoErr(err) // Could not add new CD to warehouse

	is.NoErr(whse.SellSingleCD(newID)) // Could not sell single CD
}
