package main

import (
	"fmt"
	"time"
)

	func Iterate[E any](e []E) func(func(int, E) bool) {
		return func(yield func(int, E) bool) {
			for i := 0; i <= len(e)-1; i++ {
				// to really hammer home the point about memory efficiency
				// let's add a sleep here to simulate us doing some fairly
				// intense computational work
				time.Sleep(5 * time.Second)
				if !yield(i, e[i]) {
					return
				}
			}
		}
	}

func main() {

	for x := range Iterate([]int{1, 2, 3}) {
		fmt.Println(x)
	}

}