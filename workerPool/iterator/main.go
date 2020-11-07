package main

import (
	"fmt"
	"time"
)

// Acc is a struct that accepts an integer
type Acc struct {
	id int
}

func main() {

	start := time.Now()
	for i := 0; i < 100; i++ {
		process(i)
	}
	elapsed := time.Since(start)
	fmt.Printf("tasks completed and took %v\n", elapsed) // takes approx. 10 - 11 seconds
}

func process(a int){
	fmt.Printf("Start processing %v\n", a)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finished processing %v\n", a)
}