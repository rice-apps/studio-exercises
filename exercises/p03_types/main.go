package main

import (
	"fmt"
)

func main() {
	inventory := map[string]int{
		"Whole Milk": 20, // oz
		"Oat Milk":   5,  // oz
		"Croissant":  2,  // count
	}

	// Example 1: Success
	fog := &Drink{Name: "London Fog", Size: SizeLarge, BasePrice: 4.50, Milk: MilkWhole, MilkOunces: 12}
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
	smoothie := &Drink{Name: "Mango Peach Smoothie", Size: SizeSmall, BasePrice: 5.00, Milk: MilkOat, MilkOunces: 12}
	ticket2 := Order{TicketID: "#002", Items: []MenuItem{smoothie}}
	total2, err2 := ticket2.Fulfill(inventory)
	if err2 != nil {
		fmt.Printf("Ticket %s failed: %v\n", ticket2.TicketID, err2)
	} else {
		fmt.Printf("Ticket %s fulfilled! Total: $%.2f\n", ticket2.TicketID, total2)
	}
}
