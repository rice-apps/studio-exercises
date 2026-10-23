package main

// MenuItem interface defines anything that can be added to an order.
type MenuItem interface {
	Price() float64
	Prepare(inventory map[string]int) error
}
