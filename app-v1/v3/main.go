package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 1233
}
func (i *InMemoryPlayerStore) RecordWin(name string) {}
func main() {
	handler := &PlayerServer{&InMemoryPlayerStore{}, map[string]int{}}
	log.Fatal(http.ListenAndServe(":5000", handler))
}
