package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Megan")
	want := "Hello Megan!"

	if got != want {
		t.Errorf("Greet(\"Megan\") = %q; want %q", got, want)
	}
}
