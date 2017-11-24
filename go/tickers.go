package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers use a similar mechanism to timers: a channel that is sent values.
	ticker1 := time.NewTicker(time.Millisecond * 500)

	// Here we’ll use the range builtin on the channel to iterate over the values as they arrive every 500ms.
	go func() {
		for t := range ticker1.C {
			fmt.Println("Ticker ", t)
		}
	}()


	// Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel.
	time.Sleep(time.Millisecond * 1600)

	ticker1.Stop()
	fmt.Println("Ticker stopped")

}