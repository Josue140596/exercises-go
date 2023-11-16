package contexts

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func simulateLongRunningTask(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing some work", i)
		// Done() is a channel that is closed when the context is cancelled
		case <-ctx.Done():
			fmt.Println("Cancelling because of limit time", i)
			return
		}
	}
	fmt.Println("Finished long running task")
}

func ContextsExamplesWithTimeOut() {

	// Context in Go is a powerful tool to handle cancellation operation,
	// limit execution time and metadata propagation within goroutines

	// This context is the root or father context
	rootCtx := context.Background()

	// This context is the child of the root context
	// After 3 seconds, this context will be cancelled
	childCtx, _ := context.WithTimeout(rootCtx, 3*time.Second)

	// What happens if we want only wait for 3 seconds?
	// Need to use context package
	go simulateLongRunningTask(childCtx)

	// Wait for 6 seconds to see if the long running task prints anything
	time.Sleep(6 * time.Second)
}
