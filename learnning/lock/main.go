package main

import (
	"fmt"
	"sync"
)

var count = 0

var wg sync.WaitGroup
var mutex sync.Mutex

func add() {
	mutex.Lock()
	count++
	mutex.Unlock()
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
