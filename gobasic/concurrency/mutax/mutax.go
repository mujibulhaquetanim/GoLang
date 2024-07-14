package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		fmt.Println("one appended")
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		fmt.Println("two appended")
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		fmt.Println("three appended")
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
