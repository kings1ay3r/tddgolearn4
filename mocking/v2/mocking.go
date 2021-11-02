package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buffer := &bytes.Buffer{}

	Countdown(buffer)
}

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "GO!")

}
