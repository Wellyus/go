package sync_

import (
	"fmt"
	"sync"
)

var lock_ sync.Mutex
var cond_ = sync.NewCond(&lock_)
var ok = make([]bool, 5)

func eat_init() {
	for i := range ok {
		ok[i] = true
	}
}

// return the number of left chopstick of phylosi id
func left(id int) int {
	return id
}

// return the number of right chopstick of phylosi id
func right(id int) int {
	return (id + 1) % 5
}

func eat(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// exclusion
		cond_.L.Lock()
		for !(ok[left(id)] && ok[right(id)]) {
			cond_.Wait()
		}
		//change condition, other goroutines always could not get chop1 or chop 2
		//before myself changes the state of chop1 and chop2
		//others but might could get other chops excepte these two
		ok[left(id)] = false  //chop1
		ok[right(id)] = false //chop2
		// allow other goroutines to come into critical section
		cond_.L.Unlock()

		//if I didn't change the state of condition, I couldn't unlock before my eating
		//cause I may couldn't have chopsticks anymore
		//eating,
		fmt.Printf("I'm %d phylosiphor, I'm eating now!\n", id)

		//after eating I wanna put down chopsitcks I' have used and signal to other goroutines
		cond_.L.Lock()
		ok[left(id)] = true
		ok[right(id)] = true
		//to get it!
		cond_.Broadcast()
		cond_.L.Unlock()
	}
}

func Phy_eat() {
	eat_init()
	// wait for 5 goroutines coming back
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go eat(i, &wg)
	}
	// wait for 5 goroutines coming back
	wg.Wait()
}
