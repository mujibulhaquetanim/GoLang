package main

import (
	"fmt"
	"iter"
)

// iter.seq used instead of func(func(int) bool) syntax
func CountDown(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := v; i > 0; i-- {
			if !yield(i) {
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
