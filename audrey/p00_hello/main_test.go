package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Audrey")
	want := "Hello Audrey!"

	if got != want {
		t.Errorf("Greet(\"Audrey\") = %q; want %q", got, want)
	}
}
