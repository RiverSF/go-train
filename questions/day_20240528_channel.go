package questions

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("--for")
		select {
		case c <- x:
			fmt.Println("x=", x)
			x, y = y, x+y
			fmt.Println("--x")
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Q20240528() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("--c")
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
