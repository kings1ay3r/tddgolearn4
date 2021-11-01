package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertMessage(t, got, want)
	})
	t.Run("Saying Hello World on empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
	t.Run("Saying Hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "es")
		want := "Hola, Elodie"
		assertMessage(t, got, want)
	})
	t.Run("Saying Hello in French", func(t *testing.T) {
		got := Hello("Marcus", "fr")
		want := "Bonjour, Marcus"
		assertMessage(t, got, want)
	})
}
