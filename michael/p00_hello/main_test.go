package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Michael")
	want := "Hello Michael!"

	if got != want {
		t.Errorf("Greet(\"Michael\") = %q; want %q", got, want)
	}
}
