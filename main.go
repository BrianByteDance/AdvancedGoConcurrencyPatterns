package AdvancedGoConcurrencyPatterns

import "fmt"

func demo1() {
	go f()
	go g(1, 2)
}

func g(x, y int) {}

func f() {}

func demo2() {
	c := make(chan int)
	go func() { c <- 3 }()
	n := <-c

	fmt.Print(n)
}
