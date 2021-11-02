package main

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	// buffer := &bytes.Buffer{}
	spySleeper := &SpyCountdownOperations{}
	Countdown(spySleeper, spySleeper)
	want := []string{
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}
	if !reflect.DeepEqual(want, spySleeper.Calls) {
		t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
	}
}
