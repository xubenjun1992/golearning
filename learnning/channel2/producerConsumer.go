package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producer(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	wg.Done()
}

func consumer(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}

func main() {
	c := make(chan int, 10)
	wg.Add(2)
	go producer(c)
	go consumer(c)
	wg.Wait()
}
