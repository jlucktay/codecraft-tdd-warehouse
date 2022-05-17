//go:generate mockgen -source=chart.go -destination=./chart_mock_test.go -package=warehouse_test

package warehouse

// > When someone buys a CD, we need to notify the charts of the sale, telling them the artist and title of the CD bought, and how many copies were sold.

type Chart interface {
	Notify(artist, title string, copies uint)
}
