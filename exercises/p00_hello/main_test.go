package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Andrew")
	want := "Hello Andrew!"

	if got != want {
		t.Errorf("Greet(\"Andrew\") = %q; want %q", got, want)
	}
}
