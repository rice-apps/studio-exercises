package main

import (
	"fmt"
)

// CalculateWeeklyRevenue accepts a fixed-size array of 7 daily revenue totals.
// Returns (totalRevenue, averageDailyRevenue).
func CalculateWeeklyRevenue(sales [7]float64) (float64, float64) {
	// TODO: Implement using a range loop over the array
	return 0.0, 0.0
}

// FilterOnlineOrders filters a slice of raw order logs.
// Returns a new slice containing only the logs that begin with "ONLINE:".
func FilterOnlineOrders(logs []string) []string {
	// TODO: Implement using strings.HasPrefix and append()
	return nil
}

// ParseModifiers extracts modifiers from a comma-separated ticket string.
// Example: "London Fog, extra foam,  oat milk "
// Returns: ["extra foam", "oat milk"] (The drink name itself is discarded).
func ParseModifiers(ticket string) []string {
	// TODO: Implement using strings.Split and strings.TrimSpace
	// hi
	return nil
}

// TallyDrinks creates a frequency map of drink orders.
func TallyDrinks(drinks []string) map[string]int {
	// TODO: Implement using make(map[string]int) and a range loop
	return nil
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
