package main

func star(n int, c chan int) {
	for i := 1; i < n; i++ {
		c <- i*2
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go star(cap(c), c)
	for i := range c {
		for j := 0;j<i;j++{
			print("*")
		}
		print("\n")
	}
}

