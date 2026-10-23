package main

import (
	"testing"
)

func TestOrder_Fulfill_Success(t *testing.T) {
	inv := map[string]int{
		"Whole Milk": 20,
		"Croissant":  2,
	}

	fog := &Drink{Name: "London Fog", Size: SizeLarge, BasePrice: 4.50, Milk: MilkWhole, MilkOunces: 12}
	croissant := Pastry{Name: "Croissant", BasePrice: 3.00}

	o := Order{
		TicketID: "123",
		Items:    []MenuItem{fog, croissant},
	}

	gotTotal, err := o.Fulfill(inv)
	if err != nil {
		t.Fatalf("Fulfill() returned unexpected error: %v", err)
	}

	wantTotal := 8.50 // 4.50 + 1.00 (large) + 3.00
	if !floatEqual(gotTotal, wantTotal) {
		t.Errorf("Fulfill() total = %.2f, want %.2f", gotTotal, wantTotal)
	}

	if got := inv["Whole Milk"]; got != 8 {
		t.Errorf("inventory[Whole Milk] = %d, want 8", got)
	}
	if got := inv["Croissant"]; got != 1 {
		t.Errorf("inventory[Croissant] = %d, want 1", got)
	}
}

// TODO: STUDENT ASSIGNMENT
// Write a test for the out-of-stock scenario in TestOrder_OutOfStock below!
//
// Requirements:
// 1. Setup an inventory map with only 5 oz of "Oat Milk".
// 2. Create an Order containing a Drink that requires 12 oz of MilkOat (e.g. Size SizeLarge or SizeSmall).
// 3. Call Fulfill() and verify that a non-nil error is returned.
// 4. Verify the error is of type *OutOfStockError using errors.As.
// 5. Verify the fields Item, Needed, and Available on the custom error match your setup exactly.
//
// Note: You may need to add imports at the top of this file!

func TestOrder_OutOfStock(t *testing.T) {
	t.Fatal("TODO: Implement TestOrder_OutOfStock")
}
