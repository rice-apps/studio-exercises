package main

import (
	"errors"
	"math"
	"testing"
)

const floatTolerance = 0.001

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= floatTolerance
}

func TestDrink_Price(t *testing.T) {
	tests := []struct {
		name  string
		drink *Drink
		want  float64
	}{
		{
			name:  "small drink returns base price",
			drink: &Drink{Name: "Latte", Size: SizeSmall, BasePrice: 4.00},
			want:  4.00,
		},
		{
			name:  "large drink adds size surcharge",
			drink: &Drink{Name: "Latte", Size: SizeLarge, BasePrice: 4.00},
			want:  5.00,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.drink.Price()
			if !floatEqual(got, tc.want) {
				t.Errorf("Price() = %.2f, want %.2f", got, tc.want)
			}
		})
	}
}

func TestDrink_Prepare(t *testing.T) {
	t.Run("none milk skips inventory check", func(t *testing.T) {
		drink := &Drink{Name: "Espresso", Size: SizeSmall, BasePrice: 3.00, Milk: MilkNone, MilkOunces: 0}
		inv := map[string]int{"Whole Milk": 10}

		if err := drink.Prepare(inv); err != nil {
			t.Fatalf("Prepare() returned unexpected error: %v", err)
		}
		if got := inv["Whole Milk"]; got != 10 {
			t.Errorf("inventory changed unexpectedly: got %d, want 10", got)
		}
	})

	t.Run("sufficient inventory deducts milk ounces", func(t *testing.T) {
		drink := &Drink{Name: "Cortado", Size: SizeSmall, BasePrice: 3.75, Milk: MilkWhole, MilkOunces: 4}
		inv := map[string]int{"Whole Milk": 10}

		if err := drink.Prepare(inv); err != nil {
			t.Fatalf("Prepare() returned unexpected error: %v", err)
		}
		if got := inv["Whole Milk"]; got != 6 {
			t.Errorf("inventory[Whole Milk] = %d, want 6", got)
		}
	})

	t.Run("insufficient inventory returns OutOfStockError", func(t *testing.T) {
		drink := &Drink{Name: "Large Latte", Size: SizeLarge, BasePrice: 5.00, Milk: MilkWhole, MilkOunces: 10}
		inv := map[string]int{"Whole Milk": 4}

		err := drink.Prepare(inv)
		if err == nil {
			t.Fatal("Prepare() returned nil error for shortage, want non-nil")
		}

		var oosErr *OutOfStockError
		if !errors.As(err, &oosErr) {
			t.Fatalf("expected *OutOfStockError, got %T (%v)", err, err)
		}
		if oosErr.Item != "Whole Milk" || oosErr.Needed != 10 || oosErr.Available != 4 {
			t.Errorf("OutOfStockError fields mismatch: got %+v, want Item: Whole Milk, Needed: 10, Available: 4", oosErr)
		}
	})
}
