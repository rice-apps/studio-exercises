package main

import (
	"errors"
	"testing"
)

func TestPastry_Price(t *testing.T) {
	pastry := Pastry{Name: "Scone", BasePrice: 3.25}
	if got, want := pastry.Price(), 3.25; !floatEqual(got, want) {
		t.Errorf("Price() = %.2f, want %.2f", got, want)
	}
}

func TestPastry_Prepare(t *testing.T) {
	t.Run("sufficient inventory deducts 1 unit", func(t *testing.T) {
		inv := map[string]int{"Scone": 1}
		pastry := Pastry{Name: "Scone", BasePrice: 3.25}

		if err := pastry.Prepare(inv); err != nil {
			t.Fatalf("Prepare() returned unexpected error: %v", err)
		}
		if got := inv["Scone"]; got != 0 {
			t.Errorf("inventory[Scone] = %d, want 0", got)
		}
	})

	t.Run("insufficient inventory returns OutOfStockError", func(t *testing.T) {
		inv := map[string]int{"Scone": 0}
		pastry := Pastry{Name: "Scone", BasePrice: 3.25}

		err := pastry.Prepare(inv)
		if err == nil {
			t.Fatal("Prepare() returned nil error for shortage, want non-nil")
		}

		var oosErr *OutOfStockError
		if !errors.As(err, &oosErr) {
			t.Fatalf("expected *OutOfStockError, got %T (%v)", err, err)
		}
		if oosErr.Item != "Scone" || oosErr.Needed != 1 || oosErr.Available != 0 {
			t.Errorf("OutOfStockError fields mismatch: got %+v, want Item: Scone, Needed: 1, Available: 0", oosErr)
		}
	})
}
