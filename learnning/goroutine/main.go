package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func fn1() {
	fmt.Printf("1-10的奇数:")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf(strconv.Itoa(i))
		}
	}
	fmt.Println()
	wg.Done()
}

func fn2() {
	fmt.Printf("2-10的偶数:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf(strconv.Itoa(i))
		}
	}
	fmt.Println()
	wg.Done()
}

func main() {
	wg.Add(2)
	go fn1()
	go fn2()
	wg.Wait()
}
