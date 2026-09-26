package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Claire")
	want := "Hello Claire!"

	if got != want {
		t.Errorf("Greet(\"Claire\") = %q; want %q", got, want)
	}
}
