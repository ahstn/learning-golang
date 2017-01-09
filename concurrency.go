package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	/* Goroutines
	   - Lighter weight than OS Threads
	   - Go manages goroutines (making it simpler for programmers
	   - Less switching and spinning up threads
	   - Faster startup than OS Threads
	   - Safe communication between goroutines (using channels)
	   - Uses an Actor model via Communicating Sequential Processes (CSP) */

	// GOMAXPROCS() tells go we have 2 threads and therefore enables parallelism
	// Parallelism != Concurrency. Cucurrency is independently executing processes
	runtime.GOMAXPROCS(2)

	// Adding WaitGroup is a way to tell the main() when the goroutines are finished
	// Otherwise it may exit before they finish executing
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// "go func" syntax spawns a new goroutine for that function
	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	// While the above func is waiting, this function will get switched onto the thread
	go func() {
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()

	/* Channel Types
	   Unbuffered
	       - myChannel := make(chan int)
	       - Size not specified on declaration
	       - Goroutine adding data to channel must wait until another routine picks up the data
	       - synchronous behaviour
	   Buffered
	       - myChannel := make(chan int, 5)
	       - Size is specified on declaratio
	       - Gorountines don't have to wait after adding data to the channel (unless the channel is full) */
}
