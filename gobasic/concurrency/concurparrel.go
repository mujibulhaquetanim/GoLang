package main

import (
	"fmt"
	"time"
)

func ConcurParrel() {
	start := time.Now()
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("time taken", time.Since(start))
}

func main() {
	ConcurParrel()
}

// go run -race concurparrel.go