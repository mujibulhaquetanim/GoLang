package main

import "fmt"

// unbuffered channel is an synchronous channel that is used to send data without any limit. It is used in cases where the number of messages is not fixed and it is not a requirement to send all the messages in batches.

func main() {
	charChannel := make(chan string)
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
