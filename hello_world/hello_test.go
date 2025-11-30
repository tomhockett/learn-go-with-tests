package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tom")
	want := "Hello, Tom!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
