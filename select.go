package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
//			fmt.Println(x)	
		case <- quit:
			fmt.Println("quit")//okey exited
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c,i)
		}
		//<-quit
		quit <- 5//exiting command
	}()
	fibonacci(c,quit)

}
