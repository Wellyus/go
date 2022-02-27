package sync_

import (
	"fmt"
	"sync"
)

const (
	_ int = iota
	A
	B
	C
	D
	E
	F
)

type edge struct {
	from int
	char byte
	to   int
}

var graph = []edge{
	{A, '<', B},
	{B, '>', C},
	{C, '<', D},
	{D, '_', A},
	{A, '>', E},
	{E, '<', F},
	{F, '>', D},
}
var current int = 1
var lock = sync.Mutex{}
var cond = sync.NewCond(&lock)
var wg sync.WaitGroup

func cal(current int, char byte) int {
	for i := 0; i < len(graph); i++ {
		if graph[i].from == current && graph[i].char == char {
			return graph[i].to
		}
	}
	return 0
}

func print(char byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		cond.L.Lock()
		for cal(current, char) == 0 {
			cond.Wait()
		}
		fmt.Print(string(char))
		current = cal(current, char)
		cond.Broadcast()
		cond.L.Unlock()
	}
}

func Fish() {
	var str = []byte("<<<<>>>>____")
	for _, char := range str {
		wg.Add(1)
		go print(char, &wg)
	}
	wg.Wait()
}
