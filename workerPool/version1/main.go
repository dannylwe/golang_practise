package main

import (
	"fmt"
	"sync"
	"time"
)

// Acc is a struct that accepts an integer
type Acc struct {
	id int
}

func main() {

	start := time.Now()
	channel := make(chan Acc, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for a := range channel {
			
			wg.Add(1)
			go func(){
				defer wg.Done()
				process(a)
			}()
		}
	}()

	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Nanosecond)
		channel <- Acc{id: i}
	}
	close(channel)

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("tasks completed and took %v\n", elapsed) // takes approx. 10 - 11 seconds
}

func process(a Acc) {
	fmt.Printf("Start processing %v\n", a)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finished processing %v\n", a)
}
