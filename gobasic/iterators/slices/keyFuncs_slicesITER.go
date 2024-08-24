package main

import (
	"fmt"
	"slices"
)

func main() {

	slice := []int{1, 2, 3}
	// All returns an iterator over index-value pairs in the slice
	// in the usual order.
	for x, y := range slices.All(slice) {
		fmt.Println("Returning key-value pairs: ", x, y)
	}
	// Backward returns an iterator over index-value pairs in the slice,
	// traversing it backward with descending indices.
	for x, y := range slices.Backward(slice) {
		fmt.Println("Returning backward key-value pairs: ", x, y)
	}

	//values returns a iterator over slice's values.
	for x := range slices.Values(slice) {
		fmt.Println("values: ", x)
	}

	//Collect collects values from seq into a new slice and returns it.
	newSlice := slices.Collect(slices.Values(slice))
	fmt.Println("newSlice: ", newSlice)
}
