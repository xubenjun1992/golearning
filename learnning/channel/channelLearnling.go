package main

import (
	"fmt"
	"sync"
)

func readChannel(c <-chan int) {
	for i := 1; i <= 10; i++ {
		v := <-c
		fmt.Println(v)
	}

	wg.Done()
}

func writeChannel(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	c := make(chan int, 10)
	wg.Add(2)
	go writeChannel(c)
	go readChannel(c)
	wg.Wait()
}
