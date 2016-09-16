package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "test 1"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "test 2"
	}()

	select {
	case res := <-ch:
		fmt.Println("Res 1: ", res)
	case <-time.After(time.Second * 1):
		fmt.Println("Res 1: Timeout!")
	}

	select {
	case res := <-ch2:
		fmt.Println("Res 2: ", res)
	case <-time.After(time.Second * 2):
		fmt.Println("Res 2: Timeout!")
	}
}
