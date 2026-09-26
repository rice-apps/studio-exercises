package main

import (
	"math"
	"testing"
)

func TestGreetCustomer(t *testing.T) {
	got := GreetCustomer("Zach")
	want := "Welcome to Chaus, Zach!"
	if got != want {
		t.Errorf("GreetCustomer(\"Zach\") = %q; want %q", got, want)
	}
}

func TestCalculateDrinkPrice(t *testing.T) {
	tests := []struct {
		name       string
		basePrice  float64
		extraShots int
		altMilk    bool
		wantPrice  float64
	}{
		{"Normal 0 price", 0.0, 0, false, 0.0},
		{"Negative base", -2.5, 1, false, 0.0},
		{"Standard Latte", 4.0, 0, false, 4.00},
		{"Latte w/ Alt Milk", 4.0, 0, true, 4.75},     // 4.00 + 0.75 = 4.75
		{"Latte w/ Extra Shots", 4.0, 2, false, 5.00}, // 4.00 + (2 * 0.50) = 5.00
		{"Fully Loaded", 4.0, 2, true, 5.75},          // 4.00 + 1.00 + 0.75 = 5.75
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateDrinkPrice(tt.basePrice, tt.extraShots, tt.altMilk)
			if math.Abs(got-tt.wantPrice) > 0.001 {
				t.Errorf("CalculateDrinkPrice(%.2f, %d, %v) = %.2f; want %.2f", tt.basePrice, tt.extraShots, tt.altMilk, got, tt.wantPrice)
			}
		})
	}
}

func TestApplyDiscount(t *testing.T) {
	tests := []struct {
		total        float64
		discountType string
		wantTotal    float64
		wantApplied  bool
	}{
		{5.00, "BYOM", 4.75, true},
		{5.00, "Student", 4.50, true},
		{5.00, "Faculty", 4.75, true},
		{5.00, "None", 5.00, false},
		{5.00, "UNKNOWN", 5.00, false},
	}

	for _, tt := range tests {
		t.Run(tt.discountType, func(t *testing.T) {
			gotTotal, gotApplied := ApplyDiscount(tt.total, tt.discountType)
			if math.Abs(gotTotal-tt.wantTotal) > 0.001 || gotApplied != tt.wantApplied {
				t.Errorf("ApplyDiscount(%.2f, %q) = (%.2f, %v); want (%.2f, %v)",
					tt.total, tt.discountType, gotTotal, gotApplied, tt.wantTotal, tt.wantApplied)
			}
		})
	}
}

func TestValidatePromoCode(t *testing.T) {
	tests := []struct {
		code string
		want bool
	}{
		{"CHAUS24", true},        // Valid: 7 chars, upper + digit
		{"chaus24", false},       // Invalid: no uppercase
		{"CHAUS!!", false},       // Invalid: no digit
		{"CH24", false},          // Invalid: too short (< 5)
		{"WELCOME2CHAUS", false}, // Invalid: too long (> 10)
		{"RICE2026", true},       // Valid: 8 chars, upper + digit
		{"COFFEE", false},        // Invalid: no digit
		{"12345", false},         // Invalid: no uppercase
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			got := ValidatePromoCode(tt.code)
			if got != tt.want {
				t.Errorf("ValidatePromoCode(%q) = %v; want %v", tt.code, got, tt.want)
			}
		})
	}
}
