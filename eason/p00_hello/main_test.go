package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Eason")
	want := "Hello Eason!"

	if got != want {
		t.Errorf("Greet(\"Eason\") = %q; want %q", got, want)
	}
}
