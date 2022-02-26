package sync_

import (
	"fmt"
	"sync"
)

func Mutex() {
	var count int
	var lock sync.Mutex
	increment := func(lock *sync.Mutex) {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("count is incrementing to %d\n", count)
	}

	decrement := func(lock *sync.Mutex) {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("count is decrementing to: %d\n", count)
	}

	var wg sync.WaitGroup
	// increment and decrement count in 10 goroutines
	for i := 0; i <= 4; i++ {
		wg.Add(2)
		go func(wg *sync.WaitGroup, lock *sync.Mutex) {
			defer wg.Done()
			increment(lock)
		}(&wg, &lock)
		go func(wg *sync.WaitGroup, lock *sync.Mutex) {
			defer wg.Done()
			decrement(lock)
		}(&wg, &lock)
	}
	//wait for 10 goroutines coming back!
	wg.Wait()
	fmt.Println("wg complete, and count is ", count)
}
