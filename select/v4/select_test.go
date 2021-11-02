package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got, _ := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
func TestTimeout(t *testing.T) {

	slowServer := makeDelayedServer(4 * time.Second)
	fastServer := makeDelayedServer(3 * time.Second)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := "Some Err"
	_, got := Racer(slowUrl, fastUrl)

	if got == nil {
		t.Errorf("got %q want %q", got, want)
	}
}
