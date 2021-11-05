package main

import (
	"testing"
)

func TestFileSystemStorePlayerScore(t *testing.T) {
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")

		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
