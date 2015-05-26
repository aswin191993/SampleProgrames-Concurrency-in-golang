package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{0,1,2,3,4,5,6,7}
	c := make(chan int)
	go sum(a[:len(a)-1], c)
	go sum(a[:len(a)-3], c)
	x:= <-c
	y:= <-c // receive from c
	fmt.Println(x, y)
}
