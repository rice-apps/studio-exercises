package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Olivia")
	want := "Hello Olivia!"

	if got != want {
		t.Errorf("Greet(\"Olivia\") = %q; want %q", got, want)
	}
}
