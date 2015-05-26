package main


var c = make(chan int)
var a string
var b int = 2


func f() {
	a = "\nhello, world\n"
	print(<-c)
}

func fun(){
	b=5
}

func main() {
	go f()//function will access 
	c <- 1
	go fun()// function will not access. execution of channel exited.
	print(a)
	print(b)//the value of "b" not changed.
}
