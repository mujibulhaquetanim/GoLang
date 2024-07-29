package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 5
	fmt.Println(reflect.TypeOf(a))

	var b string = "hello"
	fmt.Println(reflect.TypeOf(b))

	var c bool = true
	fmt.Println(reflect.TypeOf(c))

	var d float64 = 5.5
	fmt.Println(reflect.TypeOf(d))

}