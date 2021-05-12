package main

import "testing"

func TestHello(t *testing.T) {
	assertEquals := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to Tadashera", func(t *testing.T) {
		got := Hello("Tadashera", "portuguese")
		want := "Olá, Tadashera"

		assertEquals(t, got, want)
	})

	t.Run("say 'Hola, Tadashera'", func(t *testing.T) {
		got := Hello("Tadashera", "spanish")
		want := "Hola, Tadashera"

		assertEquals(t, got, want)
	})

	t.Run("say 'Hello, Tadashera' when an language is empty", func(t *testing.T) {
		got := Hello("Tadashera", "")
		want := "Hello, Tadashera"

		assertEquals(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"

		assertEquals(t, got, want)
	})
}
