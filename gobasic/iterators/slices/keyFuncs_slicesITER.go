package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	s1 := []Person{{"Emma", 34}, {"Millie", 20}}
	compare := func(p1, p2 Person) int {
		return cmp.Compare(p1.Age, p2.Age)
	}

	slice := []int{1, 2, 3}
	chunked := slices.Chunk(slice, 2)
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

	//AppendSeq appends values from an iterator to an existing slice
	appendedSlice := slices.AppendSeq(slice, slices.Values(newSlice))
	fmt.Println("appendedSlice: ", appendedSlice)

	//Sorted collects values from an iterator into new slice then sorts the slice
	sortedSlice := slices.Sorted(slices.Values(slice))
	fmt.Println("sortedSlice: ", sortedSlice)

	//SortedFunc is like Sorted but takes a comparison function
	sortedFunc := slices.SortedFunc(slices.Values(s1), compare)
	fmt.Println("sortedFunc: ", sortedFunc)

	//chunk returns a iterator over consecutive sub-slices of up to n elements of the slice
	for x := range chunked {
		fmt.Println("chunk: ", x)
	}
}
