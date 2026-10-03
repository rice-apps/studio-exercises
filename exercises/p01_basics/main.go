package main

import (
	"fmt"
)

// Constants for pricing
const (
	ShotPrice    float64 = 0.50
	AltMilkPrice float64 = 0.75
)

// GreetCustomer returns a formatted greeting string for the customer.
// Output format: "Welcome to Chaus, <name>!"
func GreetCustomer(name string) string {
	return fmt.Sprintf("Welcome to Chaus, %s!", name)
}

// CalculateDrinkPrice calculates the total price of a drink based on modifiers.
// Formula: basePrice + (extraShots * ShotPrice) + (AltMilkPrice if altMilk else 0.0)
// If basePrice <= 0, return 0.0.
func CalculateDrinkPrice(basePrice float64, extraShots int, altMilk bool) float64 {
	if basePrice <= 0 {
		return 0.0
	}
	total := basePrice + (float64(extraShots) * ShotPrice)
	if altMilk {
		total += AltMilkPrice
	}
	return total
}

// ApplyDiscount applies a discount code to the total using a switch statement.
// Returns (newTotal, isApplied).
// Discount rates:
// - "BYOM": minus $0.25, true
// - "Student": 10% off (multiply by 0.90), true
// - "Faculty": 5% off (multiply by 0.95), true
// - Default: no change, false
func ApplyDiscount(total float64, discountType string) (float64, bool) {
	switch discountType {
	case "BYOM":
		return total - 0.25, true
	case "Student":
		return total * 0.90, true
	case "Faculty":
		return total * 0.95, true
	default:
		return total, false
	}
}

// ValidatePromoCode checks if a Chaus discount code is valid.
// A valid promo code must:
// 1. Be between 5 and 10 characters long (inclusive).
// 2. Contain at least one numeric digit ('0' through '9').
// 3. Contain at least one uppercase letter ('A' through 'Z').
func ValidatePromoCode(code string) bool {
	if len(code) < 5 || len(code) > 10 {
		return false
	}
	hasDigit := false
	hasUpper := false
	for i := 0; i < len(code); i++ {
		if code[i] >= '0' && code[i] <= '9' {
			hasDigit = true
		}
		if code[i] >= 'A' && code[i] <= 'Z' {
			hasUpper = true
		}
	}
	return hasDigit && hasUpper
}

func main() {
	fmt.Println(GreetCustomer("Ann"))

	// Calculate: Cup of Ambition with 2 extra shots and whole milk (no alt milk upcharge)
	price1 := CalculateDrinkPrice(4.50, 2, false)
	fmt.Printf("Cup of Ambition (2 extra shots, whole milk): $%.2f\n", price1)

	// Calculate: London Fog with alt milk
	price2 := CalculateDrinkPrice(4.00, 0, true)
	fmt.Printf("London Fog (extra foam, oat milk): $%.2f\n", price2)

	finalPrice, applied := ApplyDiscount(price1, "BYOM")
	fmt.Printf("Final price for Cup of Ambition (BYOM): $%.2f (Discount applied: %v)\n", finalPrice, applied)

	isValid := ValidatePromoCode("CHAUS24")
	fmt.Printf("Is 'CHAUS24' a valid promo code? %v\n", isValid)
}
