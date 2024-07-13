package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	wg.Wait()
}
