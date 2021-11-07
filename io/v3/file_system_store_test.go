package main

import (
	"testing"
)

func TestFileSystemStorePlayerScore(t *testing.T) {
	t.Run("get player score", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFile()
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")

		want := 33

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
