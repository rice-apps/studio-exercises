package main

import (
	"errors"
	"math"
	"testing"
)

func TestOrder_Fulfill_Success(t *testing.T) {
	inv := map[string]int{
		"Whole Milk": 20,
		"Croissant":  2,
	}

	fog := &Drink{Name: "London Fog", BasePrice: 4.50, MilkType: "Whole Milk", MilkOunces: 12}
	croissant := Pastry{Name: "Croissant", BasePrice: 3.00}

	o := Order{
		TicketID: "123",
		Items:    []MenuItem{fog, croissant},
	}

	gotTotal, err := o.Fulfill(inv)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantTotal := 7.50
	if math.Abs(gotTotal-wantTotal) > 0.001 {
		t.Errorf("o.Fulfill() total = %.2f; want %.2f", gotTotal, wantTotal)
	}

	if inv["Whole Milk"] != 8 {
		t.Errorf("inventory was not deducted correctly; want 8 oz remaining, got %d", inv["Whole Milk"])
	}
	if inv["Croissant"] != 1 {
		t.Errorf("inventory was not deducted correctly; want 1 croissant remaining, got %d", inv["Croissant"])
	}
}

// TODO: STUDENT ASSIGNMENT
// Write a test for the out-of-stock scenario in TestOrder_OutOfStock below!
//
// Requirements:
// 1. Setup an inventory map with only 5 oz of "Oat Milk".
// 2. Create an Order containing a Drink that requires 12 oz of "Oat Milk".
// 3. Call Fulfill() and verify that a non-nil error is returned.
// 4. Verify the error is of type *OutOfStockError using errors.As.
// 5. Verify the fields Item, Needed, and Available on the custom error match your setup exactly.
//
// Note: You may need to add imports at the top of this file!

func TestOrder_OutOfStock(t *testing.T) {
	inv := map[string]int{
		"Oat Milk": 5,
	}

	smoothie := &Drink{Name: "Mango Peach Smoothie", BasePrice: 5.00, MilkType: "Oat Milk", MilkOunces: 12}
	o := Order{
		TicketID: "124",
		Items:    []MenuItem{smoothie},
	}

	total, err := o.Fulfill(inv)
	if err == nil {
		t.Fatal("expected an error, but got nil")
	}

	if total != 0.0 {
		t.Errorf("expected total to be 0.0 on error, got %.2f", total)
	}

	var outOfStockErr *OutOfStockError
	if !errors.As(err, &outOfStockErr) {
		t.Fatalf("expected error of type *OutOfStockError, got %T", err)
	}

	if outOfStockErr.Item != "Oat Milk" || outOfStockErr.Needed != 12 || outOfStockErr.Available != 5 {
		t.Errorf("unexpected error fields: %+v", outOfStockErr)
	}
}
