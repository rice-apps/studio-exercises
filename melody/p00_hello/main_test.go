package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Melody")
	want := "Hello Melody!"

	if got != want {
		t.Errorf("Greet(\"Melody\") = %q; want %q", got, want)
	}
}
