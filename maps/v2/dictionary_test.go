package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "TestMeaning"}

	t.Run("Search", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "TestMeanings"

		assertStrings(t, got, want)
	})
	t.Run("Search", func(t *testing.T) {
		_, got := dictionary.Search("test2")
		want := ErrNotFound.Error()

		assertStrings(t, got.Error(), want)
	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
