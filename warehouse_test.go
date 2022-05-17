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
	whse := warehouse.New(&StubChart{}, &dummyPayments{})
	newCD := warehouse.NewCD("Metallica", "S&M", "Vertigo Records", 14.99, startingCount)
	newID, err := whse.Add(newCD)
	is.NoErr(err) // Could not add new CD to warehouse

	// Act
	err = whse.SellSingleCD(newID)

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
	whse := warehouse.New(nil, &dummyPayments{})

	// Act
	err := whse.SellSingleCD(dummyID)

	// Assert
	is.Equal(err, warehouse.ErrCDNotFound) // Should receive specific error
}

// - batch receipt of CD stock from record label into warehouse
func TestBatchReceiptFromRecordLabel(t *testing.T) {
	// Arrange
	is := is.New(t)
	whse := warehouse.New(nil, &dummyPayments{})

	// Act
	err := whse.BatchReceipt("Vertigo Records",
		warehouse.WarehouseReceipt{"Metallica", "One", 10.0, 3},
		warehouse.WarehouseReceipt{"Metallica", "Two", 12.0, 4},
		warehouse.WarehouseReceipt{"Metallica", "Three", 12.5, 5},
	)

	// Assert
	is.NoErr(err)
}

// - look up a CD by its title
func TestLookUpCDByTitle(t *testing.T) {
	// Arrange
	is := is.New(t)
	whse := warehouse.New(nil, &dummyPayments{})
	newCD := warehouse.NewCD("Metallica", "S&M", "Vertigo Records", 14.99, 42)
	_, err := whse.Add(newCD)
	is.NoErr(err) // Could not add CD to warehouse

	// Act
	result, err := whse.FindCDByTitle("S&M")

	// Assert
	is.NoErr(err)
	is.Equal(result.GetTitle(), "S&M") // Title of result does not match
	is.Equal(result.GetPrice(), 14.99) // Price of result does not match
}

// - look up a CD by its title then return `warehouse.ErrCDNotFound` when CD not found.
func TestLookUpCDByTitleIsNotFound(t *testing.T) {
	// Arrange
	is := is.New(t)
	whse := warehouse.New(nil, &dummyPayments{})
	newCD := warehouse.NewCD("Metallica", "S&M", "Vertigo Records", 14.99, 42)
	_, err := whse.Add(newCD)
	is.NoErr(err) // Could not add CD to warehouse

	// Act
	_, err = whse.FindCDByTitle("dummy-title")

	// Assert
	is.Equal(err, warehouse.ErrCDNotFound)
}
