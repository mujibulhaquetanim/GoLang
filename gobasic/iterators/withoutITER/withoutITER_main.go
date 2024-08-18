package main

import "fmt"

// CountDown iterator function takes in a value and returns a func that takes in another func with a signature of `func(int) bool`
func CountDown(v int) func(func(int) bool) {

	return func(yield func(int) bool) {
		for i := v; i > 0; i-- {
			// we then start a for loop that iterates
			if !yield(i) {
				// we then return and finish our iterations
				return
			}
		}
	}
}

func main() {
	for x := range CountDown(5) {
		fmt.Println(x)
	}
}
