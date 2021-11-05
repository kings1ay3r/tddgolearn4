package main

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer removeFile()
		store := NewFileSystemPlayerStore(database)
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
		got = store.GetLeague()

		assertLeague(t, got, want)
	})
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer removeFile()
		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Chris")
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 34},
		}
		assertLeague(t, got, want)
	})
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer removeFile()
		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Pepper")
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
			{"Pepper", 1},
		}
		assertLeague(t, got, want)
	})
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, ``)
		defer removeFile()
		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Pepper")
		got := store.GetLeague()
		want := []Player{
			{"Pepper", 1},
		}
		assertLeague(t, got, want)
	})
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func CreateTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
