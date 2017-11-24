package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings

	time.Sleep(time.Second * 1)

	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	fmt.Println("Sending...")
	ping(pings, "test msg")

	pong(pings, pongs)
	fmt.Print("Received: ")

	fmt.Print(<-pongs)

}
