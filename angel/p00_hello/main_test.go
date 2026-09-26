package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Angel")
	want := "Hello Angel!"

	if got != want {
		t.Errorf("Greet(\"Angel\") = %q; want %q", got, want)
	}
}
