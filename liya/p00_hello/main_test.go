package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Liya")
	want := "Hello Liya!"

	if got != want {
		t.Errorf("Greet(\"Liya\") = %q; want %q", got, want)
	}
}
