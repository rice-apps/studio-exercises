package main

// MilkType represents the milk option for a drink.
type MilkType string

const (
	MilkNone  MilkType = "None"
	MilkWhole MilkType = "Whole Milk"
	MilkOat   MilkType = "Oat Milk"
)

// DrinkSize represents the size of a drink.
type DrinkSize int

const (
	SizeSmall DrinkSize = iota
	SizeLarge
)

// Drink represents a coffee or tea order.
type Drink struct {
	Name       string
	Size       DrinkSize
	BasePrice  float64
	Milk       MilkType
	MilkOunces int
}

// Price calculates the total price of the drink, including any size surcharge.
func (d *Drink) Price() float64 {
	if d.Size == SizeLarge {
		return d.BasePrice + 1.00
	}
	return d.BasePrice
}

// Prepare verifies and deducts the required milk from the inventory.
// Drinks with MilkNone require no milk. Returns OutOfStockError if inventory is insufficient.
func (d *Drink) Prepare(inventory map[string]int) error {
	if d.Milk == MilkNone {
		return nil
	}
	key := string(d.Milk)
	if inventory[key] < d.MilkOunces {
		return &OutOfStockError{
			Item:      key,
			Needed:    d.MilkOunces,
			Available: inventory[key],
		}
	}
	inventory[key] -= d.MilkOunces
	return nil
}
