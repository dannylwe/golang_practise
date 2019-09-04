package main

import "fmt"

type Rectangele struct {
	height int
	width int
}

type Shape interface {
	Area() int
}

// Area function gets area of rectangle
func (r *Rectangele) Area() int {
	return r.height * r.width
}

func main() {
	rect := Rectangele{5, 20}
	fmt.Println("Area of a rectangle using Area method is: ", rect.Area())
	var Irect Shape = &rect
	fmt.Println("Area of a rectangle using Shape interface is: ", Irect.Area())
}
