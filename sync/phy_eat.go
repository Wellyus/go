package sync_

import (
	"fmt"
	"sync"
)

// phylosiphors eating on a table
var lock_eat sync.Mutex
var cond_eat [5]sync.Cond
var chop = make(map[int]bool, 5)

func eat_init() {
	for i := range chop {
		chop[i] = true
		cond_eat[i] = sync.NewCond(&lock_eat)
	}
}

// return the number of left chopstick of phylosi id
func left(id int) int {
	return (id - 1 + 5) % 5
}

// return the number of right chopstick of phylosi id
func right(id int) int {
	return (id + 1) % 5
}

func eat(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		cond_eat[id].L.Lock()
		for chop[left(id)] == false || chop[right(id)] == false {
			cond_eat[id].Wait()
		}
		chop[left(id)] = false
		chop[right(id)] = false
		fmt.Printf("I'm %d phylosiphor and I'm eating now!\n", id)
		chop[left(id)] = true
		chop[right(id)] = true
		cond_eat[id].Broadcast()
		cond_eat[id].L.Unlock()
	}
}

func phy_eat() {
	eat_init()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go eat(i, &wg)
	}
	wg.Wait()
}
