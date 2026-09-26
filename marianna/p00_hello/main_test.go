package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Marianna")
	want := "Hello Marianna!"

	if got != want {
		t.Errorf("Greet(\"Marianna\") = %q; want %q", got, want)
	}
}
