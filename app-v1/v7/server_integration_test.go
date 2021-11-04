package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWInsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{}
	server := PlayerServer{&store, http.NewServeMux()}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponse(t, response.Body.String(), "3")
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}
