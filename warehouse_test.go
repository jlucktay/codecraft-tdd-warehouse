package warehouse_test

import (
	"testing"

	"github.com/matryer/is"

	warehouse "go.jlucktay.dev/codecraft-tdd-warehouse"
)

// - sell a single CD to a customer and take payment
func TestSellSingleCDWhenStockIsGreaterThanZero(t *testing.T) {
	is := is.New(t)

	const (
		startingCount      = 3
		finalCount    uint = 2
	)

	// Arrange
	whse := warehouse.New()
	newID := whse.NewCD( /*"S&M", "Metallica",*/ "Vertigo Records", 14.99, startingCount)

	// Act
	err := whse.SellSingleCD(newID)

	// Assert
	is.NoErr(err) // Could not sell single CD
	assertFinal := whse.QueryStockLevel(newID)
	is.Equal(assertFinal, finalCount) // Stock level did not reduce
}

// - attempt to sell a CD that doesn't exist
func TestSellCDThatDoesNotExist(t *testing.T) {
	is := is.New(t)

	const (
		dummyID = "dummyID"
	)

	// Arrange
	whse := warehouse.New()

	// Act
	err := whse.SellSingleCD(dummyID)

	// Assert
	is.Equal(err, warehouse.ErrCDNotFound) // Should receive specific error
}

// - batch receipt of CD stock from record label into warehouse
func TestBatchReceiptFromRecordLabel(t *testing.T) {
	// Arrange
	is := is.New(t)
	whse := warehouse.New()

	// Act
	err := whse.BatchReceipt("Vertigo Records",
		whse.NewStockedItem(10.0, 3),
		whse.NewStockedItem(12.0, 4),
		whse.NewStockedItem(12.5, 5),
	)

	// Assert
	is.NoErr(err)
}

// - look up a CD by its title 🚧
