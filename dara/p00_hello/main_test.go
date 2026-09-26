package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Dara")
	want := "Hello Dara!"

	if got != want {
		t.Errorf("Greet(\"Dara\") = %q; want %q", got, want)
	}
}
