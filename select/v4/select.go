package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

}

var twoSecondTimeout = 2 * time.Second

func Racer(a, b string) (winner string, error error) {
	return configurableRacer(a, b, twoSecondTimeout)
}

func configurableRacer(a, b string, timout time.Duration) (string, error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil

	case <-time.After(timout):
		return "", fmt.Errorf("Timeout")
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

/* func measureResponseTime(a string) time.Duration {

	startA := time.Now()
	http.Get(a)
	return time.Since(startA)
} */
