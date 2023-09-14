package AdvancedGoConcurrencyPatterns

import "fmt"

func main() {
	var x int
	var xc, yc chan int

	select {
	case xc <- x:
		// sent x on xc
	case y := <-yc:
		// received y from yc
		fmt.Print(y)
	}
}
