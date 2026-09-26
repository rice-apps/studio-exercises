package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Zimo")
	want := "Hello Zimo!"

	if got != want {
		t.Errorf("Greet(\"Zimo\") = %q; want %q", got, want)
	}
}
