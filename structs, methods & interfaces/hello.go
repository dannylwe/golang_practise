package main

import "fmt"

type shape interface {
	Area() int
}
type object interface {
	Volume() int
	Tally() int
}
type cube struct {
	side int
}

func(c cube) Area() int {
	return 6 * (c.side * c.side)
}

func(c cube) Volume() int {
	return c.side * c.side * c.side
}

func main(){
	testCube := cube{side: 10}
	fmt.Println(testCube.Volume())
	fmt.Println(testCube.Area())
}