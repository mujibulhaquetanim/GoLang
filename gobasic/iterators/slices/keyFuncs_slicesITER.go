package main

import (
	"fmt"
	"slices"
)

// All returns an iterator over index-value pairs in the slice
// in the usual order.

func main() {

	slice := []int{1, 2, 3}

	for x, y := range slices.All(slice) {
		fmt.Println(x, y)
	}
}
