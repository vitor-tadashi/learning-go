package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tadashera")

	got := buffer.String()
	want := "Hello, Tadashera"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
