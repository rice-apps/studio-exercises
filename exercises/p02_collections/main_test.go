package main

import (
	"math"
	"reflect"
	"testing"
)

func TestCalculateWeeklyRevenue(t *testing.T) {
	sales := [7]float64{10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0}
	gotTotal, gotAvg := CalculateWeeklyRevenue(sales)
	wantTotal := 280.0
	wantAvg := 40.0

	if math.Abs(gotTotal-wantTotal) > 0.001 || math.Abs(gotAvg-wantAvg) > 0.001 {
		t.Errorf("CalculateWeeklyRevenue(%v) = (%.2f, %.2f); want (%.2f, %.2f)",
			sales, gotTotal, gotAvg, wantTotal, wantAvg)
	}
}

func TestFilterOnlineOrders(t *testing.T) {
	input := []string{
		"ONLINE: Cup of Ambition",
		"IN-PERSON: Green Tea Freeze",
		"ONLINE:London Fog",
		"ERROR: Machine down",
		"ONLINE: Drip Coffee",
	}
	got := FilterOnlineOrders(input)
	want := []string{
		"ONLINE: Cup of Ambition",
		"ONLINE:London Fog",
		"ONLINE: Drip Coffee",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FilterOnlineOrders(...) = %q; want %q", got, want)
	}
}

func TestParseModifiers(t *testing.T) {
	tests := []struct {
		name   string
		ticket string
		want   []string
	}{
		{"Single modifier", "London Fog, extra foam", []string{"extra foam"}},
		{"Multiple modifiers", "Cup of Ambition,  whole milk ,  extra shot ", []string{"whole milk", "extra shot"}},
		{"No modifiers", "Green Tea Freeze", nil},
		{"Empty modifiers", "Latte, , ", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseModifiers(tt.ticket)
			// reflect.DeepEqual has issues comparing nil slice vs empty slice sometimes,
			// so we handle length 0 explicitly for safety.
			if len(got) == 0 && len(tt.want) == 0 {
				return // Both effectively empty
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseModifiers(%q) = %q; want %q", tt.ticket, got, tt.want)
			}
		})
	}
}

func TestTallyDrinks(t *testing.T) {
	input := []string{
		"Green Tea Freeze",
		"Cup of Ambition",
		"Green Tea Freeze",
		"London Fog",
		"London Fog",
		"London Fog",
	}
	got := TallyDrinks(input)
	want := map[string]int{
		"Green Tea Freeze": 2,
		"Cup of Ambition":  1,
		"London Fog":       3,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TallyDrinks(...) = %v; want %v", got, want)
	}
}
