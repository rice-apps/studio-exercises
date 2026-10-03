package main

import (
	"fmt"
	"strings"
)

// CalculateWeeklyRevenue accepts a fixed-size array of 7 daily revenue totals.
// Returns (totalRevenue, averageDailyRevenue).
func CalculateWeeklyRevenue(sales [7]float64) (float64, float64) {
	var total float64
	for _, daily := range sales {
		total += daily
	}
	return total, total / 7.0
}

// FilterOnlineOrders filters a slice of raw order logs.
// Returns a new slice containing only the logs that begin with "ONLINE:".
func FilterOnlineOrders(logs []string) []string {
	var onlineOrders []string
	for _, log := range logs {
		if strings.HasPrefix(log, "ONLINE:") {
			onlineOrders = append(onlineOrders, log)
		}
	}
	return onlineOrders
}

// ParseModifiers extracts modifiers from a comma-separated ticket string.
// Example: "London Fog, extra foam,  oat milk "
// Returns: ["extra foam", "oat milk"] (The drink name itself is discarded).
func ParseModifiers(ticket string) []string {
	parts := strings.Split(ticket, ",")
	if len(parts) <= 1 {
		return []string{}
	}
	var mods []string
	for _, part := range parts[1:] {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			mods = append(mods, trimmed)
		}
	}
	if mods == nil {
		return []string{}
	}
	return mods
}

// TallyDrinks creates a frequency map of drink orders.
func TallyDrinks(drinks []string) map[string]int {
	tally := make(map[string]int)
	for _, drink := range drinks {
		tally[drink]++
	}
	return tally
}

func main() {
	weeklySales := [7]float64{120.50, 200.0, 150.75, 180.0, 210.25, 300.0, 250.0}
	total, avg := CalculateWeeklyRevenue(weeklySales)
	fmt.Printf("Weekly Total: $%.2f, Daily Average: $%.2f\n\n", total, avg)

	logs := []string{
		"ONLINE: Green Tea Freeze",
		"IN-PERSON: Cup of Ambition",
		"ONLINE: London Fog",
	}
	onlineOnly := FilterOnlineOrders(logs)
	fmt.Printf("Filtered online orders: %v\n\n", onlineOnly)

	mods := ParseModifiers("London Fog, extra foam, whole milk")
	fmt.Printf("Parsed modifiers: %q\n\n", mods)

	drinkList := []string{
		"Green Tea Freeze", "Cup of Ambition", "London Fog",
		"Green Tea Freeze", "Cup of Ambition",
	}
	tallies := TallyDrinks(drinkList)
	fmt.Println("Drink tallies:")
	for drink, count := range tallies {
		fmt.Printf("- %s: %d\n", drink, count)
	}
}
