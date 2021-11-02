package main

import (
	"net/http"
	"time"
)

func main() {

}

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
func measureResponseTime(a string) time.Duration {

	startA := time.Now()
	http.Get(a)
	return time.Since(startA)
}
