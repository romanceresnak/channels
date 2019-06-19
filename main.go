package main

import (
	"fmt"
)

func main(){
	c := make(chan int)

	s := make(chan string)

	b := make(chan bool)

	f := make(chan float32)

	go send4(f)
	receive4(f)

	go send2(s)
	receive2(s)

	go send(c)
	receive(c)

	go send3(b)
	receive3(b)
}

func send(c chan<- int){
		c <- 42
}

func receive(c <-chan int){
	fmt.Println("Value from channel is : ", <-c)
}

func send2(s chan<- string){
	s <- "Ahoj"
}

func receive2(str <-chan string){
	fmt.Println("String value is : ", <-str)
}

func send3(b chan<- bool){
	b <- true
}

func receive3(b <-chan bool){
	fmt.Println("Boolean value is : ", <-b)
}

func send4(f chan<- float32){
	f <- 3.12
}

func receive4(f <-chan float32){
	fmt.Println("Float value is : ", <-f)
}