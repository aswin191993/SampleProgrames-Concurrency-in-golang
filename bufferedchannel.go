package main

import "fmt"

func main() {
	c := make(chan int, 4)
	c <- 1
	c <- 2
	c <- 3
	c <- 5
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

