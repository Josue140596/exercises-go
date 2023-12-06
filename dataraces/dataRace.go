package dataraces

import (
	"fmt"
	"sync"
)

func ExampleDataRace() {
	wg := sync.WaitGroup{}
	x := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			x++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("(ExampleDataRace) x =", x)
}
