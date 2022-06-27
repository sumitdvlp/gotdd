package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sumi")
	want := "Hello, Sumi"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
