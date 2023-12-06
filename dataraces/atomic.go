package dataraces

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func WorkingWithAtomic() {

	var wg sync.WaitGroup
	var x int32

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("(WorkingWithAtomic) x =", x)
}
