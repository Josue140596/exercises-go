package routines

import (
	"fmt"
	"math/rand"
	"time"
)

// We're going to simulate a Building process
func WithRoutines() {
	begin := time.Now()
	rand.NewSource(begin.UnixNano())

	fmt.Println("========Building begins WITH GoRoutines========")

	go worker("Bob", "Installing a door")
	go worker("Alice", "Mowing the lawn")
	go worker("John", "Painting the walls")
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("========Building ended in %.2f seconds ========\n", time.Since(begin).Seconds())
}
