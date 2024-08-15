package main

import "fmt"

// buffered channel is an asynchronous channel that is used to send data in batches.

func main() {
	charChannel := make(chan string, 4) //buffered channel with capacity of 4 elements
	chars := []string{"a", "b", "c", "d"}

	for _, v := range chars {
		charChannel <- v
		fmt.Println("sent", v)
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println("Received", result)
	}
}
