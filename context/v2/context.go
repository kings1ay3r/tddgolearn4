package main

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)

		// TODO need to understand this better
		/* ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		} */
		// store.Cancel()
		// fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}
