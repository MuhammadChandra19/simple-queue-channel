package main

import (
	"fmt"
	"time"
)

func main() {
	normalQueue := make(chan string)
	highPriorityQueue := make(chan string)

	// Goroutine to push normal priority notifications every 1 second
	go func() {
		value := 0
		for {
			time.Sleep(1 * time.Second)
			normalQueue <- fmt.Sprintf("ch2 %d", value)
			value++
		}
	}()

	// Goroutine to push high priority notifications every 10 seconds
	go func() {
		// value := 0
		for {
			time.Sleep(10 * time.Second)
			highPriorityQueue <- "ch1 x"
			// value++
		}
	}()

	// Goroutine to process notifications every 2 seconds
	go func() {
		for {

			// Process high priority notifications first
			select {
			case notif := <-highPriorityQueue:
				fmt.Println(notif)
			default:
				// If no high priority notifications, process normal priority
				select {
				case notif := <-normalQueue:
					fmt.Println(notif)
				default:
					// No notifications to process
				}
			}

			time.Sleep(2 * time.Second)
		}
	}()

	// Prevent the main function from exiting
	select {}
}
