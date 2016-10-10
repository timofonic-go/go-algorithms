package main

import (
	"fmt"
)

func main() {

	jobs := make(chan int)  // channel for jobs
	done := make(chan bool) // channel for finish signal

	// Start the worker in background.
	// Listen to jobs and process them upon receiving.
	// When all jobs are done, signal that the worker has finished.
	go func() {

		// Infinite loop to listen to incoming jobs.
		for {
			job, more := <-jobs
			// "more" value will be false if jobs has been closed and
			// all values in the channel have already been received.
			if more {
				fmt.Println("Received job ", job)
			} else {
				fmt.Println("All jobs received")

				// Signal that the worker has finished an the application can proceed.
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Send job ", i)
	}

	// I'm not going to send any more jobs.
	// I expect a worked to finish and application to continue after that.
	close(jobs)

	// Block the application while the worker is processing the jobs.
	<-done

	// Worked has finished.
	// Application has been unblocked.
	fmt.Println("All jobs complete")
}
