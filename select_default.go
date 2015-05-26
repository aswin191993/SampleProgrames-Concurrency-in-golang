package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("100ms")
		case <-boom:
			fmt.Println("exit 500ms")
			return//exit loop else infinite loop.
		default:
			fmt.Println("sleep 20ms")
			time.Sleep(20 * time.Millisecond)
		}
	}
}

