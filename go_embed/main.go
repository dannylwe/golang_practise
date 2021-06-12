package main

import (
	"fmt"
	_ "embed"
)

//go:embed sample.txt
var lorem string

func main() {
	fmt.Println("starting application")

	fmt.Println("lorem text")
	fmt.Println(lorem)
}