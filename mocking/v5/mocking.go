package main

import (
	"bytes"
	"fmt"
	"io"
)

const countdownStart = 3
const finalWord = "GO!"

func main() {
	buffer := &bytes.Buffer{}
	sleeper := &ConfigurableSleeper{}
	Countdown(buffer, sleeper)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}
