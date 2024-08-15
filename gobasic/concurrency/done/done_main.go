package main

import (
	"fmt"
	"time"
)

func doWork(done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Working...")
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	fmt.Println("Stopping...\ndone")
	close(done)
}
