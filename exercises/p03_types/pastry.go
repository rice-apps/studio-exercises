package main

// Pastry represents a bakery item.
type Pastry struct {
	Name      string
	BasePrice float64
}

// Price returns the base price of the pastry.
func (p Pastry) Price() float64 {
	// TODO: Implement Price.
	return 0.0
}

// Prepare deducts one unit of this pastry from the inventory.
// Returns OutOfStockError if the pastry is not available.
func (p Pastry) Prepare(inventory map[string]int) error {
	// TODO: Implement Prepare.
	return nil
}
