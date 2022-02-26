package sync_

import (
	"fmt"
	"time"
)

func Select() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Println(workCounter, "loops")
}
