package main

import "fmt"

func fn1(x int) {
	x = 10
}

func fn2(x *int) {
	*x = 40
}

func fn3(num *[]int) {
	for i := range *num {
		(*num)[i] *= 2
	}
}

func main() {
	num := &[]int{1, 3, 5, 7}
	fn3(num)
	fmt.Println(*num)
}
