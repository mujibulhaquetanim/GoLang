package main

import (
	"fmt"
	"maps"
)

func main() {

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"c": 3, "d": 4}
	// All returns an iterator over index-value pairs in the map
	// in the usual order.
	for k, v := range maps.All(m1) {
		fmt.Println(k, v)
	}
	//keys returns an iterator over map's keys
	for k := range maps.Keys(m1) {
		fmt.Println(k)
	}
	//values returns an iterator over map's values
	for v := range maps.Values(m1) {
		fmt.Println(v)
	}
	//insert adds the key-value pair from an iterator to an existing map {overwriting the exiting value}
	maps.Insert(m1, maps.All(m2))
	fmt.Println(m1)
	//collects collects key value pair from an iterator into a new map and returns it
	m3 := maps.Collect(maps.All(m1))
	fmt.Println(m3)
}
