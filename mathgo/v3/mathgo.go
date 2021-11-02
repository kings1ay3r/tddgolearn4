package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

type Point struct {
	X float64
	Y float64
}

func SecondHand(w io.Writer, t time.Time) Point {
	p := secondHandPoint(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
	return p
}
func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}
func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	//scale, flip, rotate
	p := Point{x * secondHandLength, y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}
func main() {
	t := time.Now()
	SVGWriter(os.Stdout, t)
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`
