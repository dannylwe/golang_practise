package shapes

import (
	"math"
)

// Reactangle struct to define a rectangle
type Rectangle struct {
	Width float64
	Length float64
}

type Circle struct {
	Radius float64
}

// Perimeter function gets the 2*length + 2*width the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Length + rectangle.Width)
}

//Area function finds the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

//Area function finds the area of a circle
func (c Circle) Area() float64{
	return math.Pi *c.Radius * c.Radius
}
