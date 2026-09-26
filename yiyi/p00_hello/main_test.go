package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Yiyi")
	want := "Hello Yiyi!"

	if got != want {
		t.Errorf("Greet(\"Yiyi\") = %q; want %q", got, want)
	}
}
