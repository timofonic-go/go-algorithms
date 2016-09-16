package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "channel 1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received #1: ", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received #2: ", msg2)
		}
	}
}
