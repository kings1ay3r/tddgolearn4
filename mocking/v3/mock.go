package main

import "time"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (*DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
