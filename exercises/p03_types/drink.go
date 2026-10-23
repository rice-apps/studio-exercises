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

// If you're interested, you can read about the iota keyword:
// https://go.dev/ref/spec#Constant_declarations
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
	// TODO: Implement Price based on drink sizing rules.
	return 0.0
}

// Prepare verifies and deducts the required milk from the inventory.
// Drinks with MilkNone require no milk. Returns OutOfStockError if inventory is insufficient.
func (d *Drink) Prepare(inventory map[string]int) error {
	// TODO: Implement Prepare.
	return nil
}
