package main

// Pastry represents a bakery item.
type Pastry struct {
	Name      string
	BasePrice float64
}

// Price returns the base price of the pastry.
func (p Pastry) Price() float64 {
	return p.BasePrice
}

// Prepare deducts one unit of this pastry from the inventory.
// Returns OutOfStockError if the pastry is not available.
func (p Pastry) Prepare(inventory map[string]int) error {
	if inventory[p.Name] < 1 {
		return &OutOfStockError{
			Item:      p.Name,
			Needed:    1,
			Available: inventory[p.Name],
		}
	}
	inventory[p.Name]--
	return nil
}
