package routines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// We're going to simulate a Building process
func WithRoutinesWaitGroup() {
	begin := time.Now()
	rand.NewSource(begin.UnixNano())

	fmt.Println("========Building begins WITH GoRoutines========")
	// Create a WaitGroup
	wg := &sync.WaitGroup{}

	go workerWithWaitGroup("Bob", "Installing a door", wg)
	wg.Add(1)
	go workerWithWaitGroup("Alice", "Mowing the lawn", wg)
	wg.Add(1)
	go workerWithWaitGroup("John", "Painting the walls", wg)
	wg.Add(1)

	wg.Wait()
	fmt.Printf("========Building ended in %.2f seconds ========\n", time.Since(begin).Seconds())
}

// Worker initiates a job and takes a random number of seconds between 3 and 5 to finish it
func workerWithWaitGroup(name, job string, wg *sync.WaitGroup) {
	fmt.Printf("Worker %s has started to work on %s\n", name, job)
	for i := 0; i < 10; i++ {
		fmt.Printf("Worker %s is working on (%d/%d)...\n", name, i+1, 10)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("Worker %s has finished working on %s\n", name, job)
	// Signal the WaitGroup that this worker is done
	wg.Done()
}
