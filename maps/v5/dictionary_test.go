package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "TestMeaning"}

	t.Run("Search", func(t *testing.T) {
		word := "test"
		meaning := "TestMeaning"

		assertMeaning(t, word, meaning, dictionary)
	})
	/* t.Run("Search", func(t *testing.T) {
		word := "test2"
		meaning := ""

		assertMeaning(t, word, meaning, dictionary)
	}) */
	t.Run("Add Existing", func(t *testing.T) {
		err := dictionary.Add("test", "TestMeaning2")
		meaning := "TestMeaning"
		word := "test"
		assertError(t, err, ErrWordExists)
		assertMeaning(t, word, meaning, dictionary)
	})
	t.Run("Add", func(t *testing.T) {
		dictionary.Add("test2", "TestMeaning2")
		meaning := "TestMeaning2"
		word := "test2"

		assertMeaning(t, word, meaning, dictionary)
	})
	t.Run("Update Existing", func(t *testing.T) {
		err := dictionary.Update("test", "TestMeaning2")
		meaning := "TestMeaning2"
		word := "test"
		assertError(t, err, nil)
		assertMeaning(t, word, meaning, dictionary)
	})
}
func assertMeaning(t testing.TB, word, want string, d Dictionary) {
	t.Helper()
	got, _ := d.Search(word)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
