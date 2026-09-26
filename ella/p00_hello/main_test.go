package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Ella")
	want := "Hello Ella!"

	if got != want {
		t.Errorf("Greet(\"Ella\") = %q; want %q", got, want)
	}
}
