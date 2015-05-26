package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
//say("first call") dealy with "1000 * time.Millisecond" time but "go say " act as multi thread function.
//So these are processing in parallel.
	go say("first call new goroutine1")//first thread
	go say("first call new goroutine2")//second thread	
	say("first call")
//But say("second call") here not calling multi thread. dealy will effected every processing.
	say("second call")
}

