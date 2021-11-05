package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func(*os.File, string) {
		<-c
		db.Close()
		// os.Remove(dbFileName)

		defer os.Exit(1)
	}(db, dbFileName)

	for {
		log.Fatal(http.ListenAndServe(":5000", NewPlayerServer(store)))
	}
}
