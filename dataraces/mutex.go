package dataraces

import (
	"fmt"
	"sync"
)

func WorkingWithMutex() {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	x := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mtx.Lock()
			x++
			mtx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("(WorkingWithMutex) x =", x)
}
