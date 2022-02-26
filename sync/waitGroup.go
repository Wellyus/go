package sync_

import (
	"fmt"
	"sync"
)

func WaitGroup() {
	var wg sync.WaitGroup
	// 2 goroutines to be waited
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("1st goroutine has done!")
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("2nd goroutines has done!")
	}(&wg)
	// wait for two goroutines completing
	wg.Wait()
	fmt.Println("All goroutines complete.")
}
