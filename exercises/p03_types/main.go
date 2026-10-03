package main

import (
	"fmt"
)

// OutOfStockError represents an inventory shortage.
type OutOfStockError struct {
	Item      string
	Needed    int
	Available int
}

// Error satisfies the error interface.
// Message format: "out of stock: need %d of %s but have %d"
func (e *OutOfStockError) Error() string {
	// TODO: Implement Error() string method
	return ""
}

// MenuItem interface defines anything that can be added to an order.
type MenuItem interface {
	Price() float64
	Prepare(inventory map[string]int) error
}

// Drink represents a coffee or tea order.
type Drink struct {
	Name       string
	BasePrice  float64
	MilkType   string
	MilkOunces int
}

// TODO: Implement MenuItem interface for Drink using pointer receivers (*Drink)

func (d *Drink) Price() float64 {
	return 0.0
}

func (d *Drink) Prepare(inventory map[string]int) error {
	return nil
}

// Pastry represents a bakery item.
type Pastry struct {
	Name      string
	BasePrice float64
}

// TODO: Implement MenuItem interface for Pastry using value receivers (Pastry)

func (p Pastry) Price() float64 {
	return 0.0
}

func (p Pastry) Prepare(inventory map[string]int) error {
	return nil
}

// Order represents a customer's total ticket.
type Order struct {
	TicketID string
	Items    []MenuItem // Interface slice can hold *Drink and Pastry!
}

// Fulfill calculates the total price and safely deducts inventory.
// - Iterates over Items, summing Price() and calling Prepare().
// - If any Prepare() fails, abort and immediately return (0.0, err).
// - On success, return (total, nil).
func (o Order) Fulfill(inventory map[string]int) (float64, error) {
	// TODO: Implement Fulfill
	return 0.0, nil
}

func main() {
	inventory := map[string]int{
		"Whole Milk": 20, // oz
		"Oat Milk":   5,  // oz
		"Croissant":  2,  // count
	}

	// Example 1: Success
	fog := &Drink{Name: "London Fog", BasePrice: 4.50, MilkType: "Whole Milk", MilkOunces: 12}
	croissant := Pastry{Name: "Croissant", BasePrice: 3.00}

	ticket1 := Order{TicketID: "#001", Items: []MenuItem{fog, croissant}}
	total1, err1 := ticket1.Fulfill(inventory)
	if err1 != nil {
		fmt.Printf("Ticket %s failed: %v\n", ticket1.TicketID, err1)
	} else {
		fmt.Printf("Ticket %s fulfilled! Total: $%.2f\n", ticket1.TicketID, total1)
		fmt.Printf("Remaining Whole Milk: %d oz\n\n", inventory["Whole Milk"])
	}

	// Example 2: Out of Stock
	smoothie := &Drink{Name: "Mango Peach Smoothie", BasePrice: 5.00, MilkType: "Oat Milk", MilkOunces: 12}
	ticket2 := Order{TicketID: "#002", Items: []MenuItem{smoothie}}
	total2, err2 := ticket2.Fulfill(inventory)
	if err2 != nil {
		fmt.Printf("Ticket %s failed: %v\n", ticket2.TicketID, err2)
	} else {
		fmt.Printf("Ticket %s fulfilled! Total: $%.2f\n", ticket2.TicketID, total2)
	}
}
