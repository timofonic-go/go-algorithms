package main

import (
	"fmt"
	"time"
)

// The first timer will expire ~2s after we start the program,
// but the second should be stopped before it has a chance to expire.
func main() {

	// Timers represent a single event in the future.
	// You tell the timer how long you want to wait,
	// and it provides a channel that will be notified at that time.
	timer1 := time.NewTimer(time.Second * 2)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.
	<-timer1.C

	fmt.Println("Time 1 expired!")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired!")
	}()

	// You can cancel the timer before it expires.
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped!")
	}
}
