package main

// Order represents a customer's total ticket.
type Order struct {
	TicketID string
	Items    []MenuItem
}

// Fulfill prepares each item in the order against inventory and calculates the total price.
// If any item fails to prepare, fulfillment halts immediately and returns 0.0 and the error.
func (o Order) Fulfill(inventory map[string]int) (float64, error) {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price()
		if err := item.Prepare(inventory); err != nil {
			return 0.0, err
		}
	}
	return total, nil
}
