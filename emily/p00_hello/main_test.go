package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Emily")
	want := "Hello Emily!"

	if got != want {
		t.Errorf("Greet(\"Emily\") = %q; want %q", got, want)
	}
}
