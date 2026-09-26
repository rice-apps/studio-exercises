package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Joya")
	want := "Hello Joya!"

	if got != want {
		t.Errorf("Greet(\"Joya\") = %q; want %q", got, want)
	}
}
