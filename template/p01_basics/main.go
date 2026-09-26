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
	// TODO: Implement this function using fmt.Sprintf
	return ""
}

// CalculateDrinkPrice calculates the total price of a drink based on modifiers.
// Formula: basePrice + (extraShots * ShotPrice) + (AltMilkPrice if altMilk else 0.0)
// If basePrice <= 0, return 0.0.
func CalculateDrinkPrice(basePrice float64, extraShots int, altMilk bool) float64 {
	// TODO: Implement this function using variables, constants, and if/else logic
	return 0.0
}

// ApplyDiscount applies a discount code to the total using a switch statement.
// Returns (newTotal, isApplied).
// Discount rates:
// - "BYOM": minus $0.25, true
// - "Student": 10% off (multiply by 0.90), true
// - "Faculty": 5% off (multiply by 0.95), true
// - Default: no change, false
func ApplyDiscount(total float64, discountType string) (float64, bool) {
	// TODO: Implement this function using a switch statement
	return 0.0, false
}

// ValidatePromoCode checks if a Chaus discount code is valid.
// A valid promo code must:
// 1. Be between 5 and 10 characters long (inclusive).
// 2. Contain at least one numeric digit ('0' through '9').
// 3. Contain at least one uppercase letter ('A' through 'Z').
func ValidatePromoCode(code string) bool {
	// TODO: Implement using len(), a for loop, and byte comparisons
	return false
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
