package sync_

import (
	"fmt"
	"sync"
)

func Cond() {
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func() {
		cond.L.Lock()
		if len(queue) > 0 {
			fmt.Println("Removed from queue ", queue[0])
			queue = queue[1:]
		}
		cond.L.Unlock()
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for len(queue) > 2 {
			cond.Wait() //解锁,挂起当前goroutine
		}
		// critical section
		fmt.Println("Adding to queue ", i*i)
		queue = append(queue, i*i)
		cond.L.Unlock()
		go removeFromQueue()
	}
}
