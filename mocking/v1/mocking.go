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
	fmt.Fprintf(out, "3")

}
