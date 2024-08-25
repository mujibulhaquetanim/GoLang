package main

import (
	"fmt"
	"math/rand"
)
//higerOrderFunc, return a chan of type of passed func.
func RepeatFunc[T any, K any](done <-chan K, fn func() T) chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		//infinite loop
		for {
			select {
			case <-done:
				return
				//accessing stream via closures
			case stream <- fn():
			}
		}
	}()

	//the goRoutine infinitely shooting fn() data to the stream which is in outside of the goRoutine yet can doing so by closure.
	return stream
}

func main() {
	done := make(chan int)
	defer close(done)

	RandNumFetcher := func() int {
		return rand.Intn(50000)
	}
	for randomValue := range RepeatFunc(done, RandNumFetcher) {
		fmt.Println(randomValue)
	}
}
