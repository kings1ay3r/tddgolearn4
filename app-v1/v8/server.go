package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {

	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.ServeHTTP))
	router.Handle("/players/", http.HandlerFunc(p.ServeHTTP))
	p.Handler = router
	return p
}
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processPost(w, r)
	case http.MethodGet:
		p.processGet(w, r)
	}
}
func (p *PlayerServer) LeagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer) processGet(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)

}
func (p *PlayerServer) processPost(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.RecordWin(player)

	leagueTable := []Player{
		{"Chris", 20},
	}
	json.NewEncoder(w).Encode(leagueTable)

	w.WriteHeader(http.StatusOK)

}
