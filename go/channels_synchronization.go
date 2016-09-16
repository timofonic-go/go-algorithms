package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a signal channel.
	ch := make(chan struct{})

	go worker(ch)

	<-ch
	fmt.Println("Done!")
}

func worker(ch chan struct{}) {
	fmt.Println("Working...")
	time.Sleep(time.Second * 2)
	ch <- struct{}{}
}
