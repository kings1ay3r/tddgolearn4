package main

import (
	"io/ioutil"
	"log"
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
		store, _ := NewFileSystemPlayerStore(database)
		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(t, got, want)
		got = store.GetLeague()

	})
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer removeFile()
		store, _ := NewFileSystemPlayerStore(database)
		store.RecordWin("Chris")
		got := store.GetLeague()
		want := []Player{
			{"Chris", 34},
			{"Cleo", 10},
		}
		AssertLeague(t, got, want)
	})
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer removeFile()
		store, _ := NewFileSystemPlayerStore(database)
		store.RecordWin("Pepper")
		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
			{"Pepper", 1},
		}
		AssertLeague(t, got, want)
	})
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, `[]`)
		defer removeFile()
		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			{
				log.Fatalf("problem creating file system player store, %v ", err)
			}

		}
		store.RecordWin("Pepper")
		got := store.GetLeague()
		want := []Player{
			{"Pepper", 1},
		}
		AssertLeague(t, got, want)
	})
	t.Run("league from empty file", func(t *testing.T) {
		database, removeFile := CreateTempFile(t, ``)
		defer removeFile()
		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			{
				log.Fatalf("problem creating file system player store, %v ", err)
			}

		}
		store.RecordWin("Pepper")
		got := store.GetLeague()
		want := []Player{
			{"Pepper", 1},
		}
		AssertLeague(t, got, want)
	})
}

func AssertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	if err != nil {

		t.Fatalf("could not create temp file %v", err)
	}
}
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	AssertNoError(t, err)

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
