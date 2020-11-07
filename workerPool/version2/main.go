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
	channel := make(chan Acc, 1) // job queue
	workerPoolSize := 100
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			
			go func(){
				defer wg.Done()
				for num := range channel {
					process(num)
				}
			}()
		}
	}()

	for i := 0; i < 100; i++ {
		channel <- Acc{id: i} //jobs
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
