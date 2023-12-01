package routines

import (
	"fmt"
	"math/rand"
	"time"
)

// We're going to simulate a Building process
func WithoutRoutines() {
	begin := time.Now()
	rand.NewSource(begin.UnixNano())

	fmt.Println("========Building begins without GoRoutines========")
	worker("Bob", "Installing a door")
	worker("Alice", "Mowing the lawn")
	worker("John", "Painting the walls")
	fmt.Printf("========Building ended in %.2f seconds ========\n", time.Since(begin).Seconds())
}

// Worker initiates a job and takes a random number of seconds between 3 and 5 to finish it
func worker(name, job string) {
	fmt.Printf("Worker %s has started to work on %s\n", name, job)
	for i := 0; i < 10; i++ {
		fmt.Printf("Worker %s is working on (%d/%d)...\n", name, i+1, 10)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("Worker %s has finished working on %s\n", name, job)
}
