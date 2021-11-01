package main

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0, name: "Rectangle"},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0, name: "Triangle"},
		{shape: Circle{Radius: 10}, want: 314.1592653589793, name: "Circle"},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%q : got %g want %g", tt.name, got, tt.want)
			}
		})
	}
}
