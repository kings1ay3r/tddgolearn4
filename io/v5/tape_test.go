package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTape_Write(t *testing.T) {
	db, _ := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer db.Close()
	defer os.Remove(dbFileName)
	// defer clean()
	tape := &tape{db}
	tape.Write([]byte("abc"))
	db.Seek(0, 0)
	newFileContents, err := ioutil.ReadAll(db)
	if err != nil {

		got := string(newFileContents)
		want := "abc"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
