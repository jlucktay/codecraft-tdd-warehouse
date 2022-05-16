package warehouse_test

import (
	"testing"

	"github.com/matryer/is"

	warehouse "go.jlucktay.dev/tdd-cd-warehouse"
)

// - sell a single CD to a customer and take payment
func TestSellSingleCD(t *testing.T) {
	is := is.New(t)

	// Arrange
	whse := warehouse.New()
	newID := whse.NewCD("S&M", "Metallica", "Vertigo Records", 14.99, 3)
	startingCount := whse.QueryStockLevel(newID)

	// Act
	whse.SellCD(newID)

	// Assert
	finalCount := whse.QueryStockLevel(newID)
	is.True(finalCount == startingCount-1)
}
