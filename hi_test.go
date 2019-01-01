package main

import "testing"

func TestHi(t *testing.T) {
	got := Hi()
	want := "hi asep"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
