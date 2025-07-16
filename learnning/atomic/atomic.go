package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var count int32 = 0

func add() {
	atomic.AddInt32(&count, 1)
	wg.Done()
}

func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(count)
}
