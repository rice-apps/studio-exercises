package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Katherine")
	want := "Hello Katherine!"

	if got != want {
		t.Errorf("Greet(\"Katherine\") = %q; want %q", got, want)
	}
}
