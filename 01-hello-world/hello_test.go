package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Arturo")
	want := "Hello, Arturo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
