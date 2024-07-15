package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("interface type: %T, value %v\n", i, i)
}

func main() {
	//empty interface
	var i interface{} = 1
	var j interface{} = "hello"
	var empty interface{}

	slice := []int{1, 2, 3}
	slice = append(slice, 4)

	describe(slice)
	describe(i)
	describe(j)
	describe(empty)

}
