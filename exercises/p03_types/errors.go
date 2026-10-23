package main

// OutOfStockError represents an inventory shortage for an ingredient or item.
type OutOfStockError struct {
	Item      string
	Needed    int
	Available int
}

// Error returns a formatted message describing the stock shortage.
// Expected format: "out of stock: need <needed> of <item> but have <available>"
// This method allows OutOfStockError to fulfill Go's error interface.
func (e *OutOfStockError) Error() string {
	// TODO: Implement Error.
	return ""
}
