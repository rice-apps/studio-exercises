package main

import "fmt"

// OutOfStockError represents an inventory shortage for an ingredient or item.
type OutOfStockError struct {
	Item      string
	Needed    int
	Available int
}

// Error returns a formatted message describing the stock shortage.
// Message format: "out of stock: need <needed> of <item> but have <available>"
func (e *OutOfStockError) Error() string {
	return fmt.Sprintf("out of stock: need %d of %s but have %d", e.Needed, e.Item, e.Available)
}
