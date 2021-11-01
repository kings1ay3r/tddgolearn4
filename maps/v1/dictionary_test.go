package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "TestMeaning"}

	got := Search(dictionary, "test")
	want := "TestMeaning"

	assertStrings(t, got, want)
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
