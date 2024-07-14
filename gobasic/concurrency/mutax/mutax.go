package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	//RWMutex added instead of Mutex, though it not been used like sync.Mutex
	//writing is the first priority in RWMutex and reading is the second, if one goroutine is comes to write the same address, all other goroutines will be blocked that are reading that address for writing.
	//it is best to use in function where it is used instead of the main variable that holds the value, i.e. instead of locking/unlocking the score, use it in the function.-> mut.Lock() var score = []int{0} mut.Unlock()
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4) //4 goroutines
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		fmt.Println("one appended")
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		fmt.Println("two appended")
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		fmt.Println("three appended")
	}(wg, mut)

	//this goroutine allowed for reading only.
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		m.RLock()
		fmt.Println("this is for RLock/unlock")
		m.RUnlock()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
