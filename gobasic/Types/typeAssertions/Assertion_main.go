package main

import "fmt"

func main() {
	// as interface can be assigned to any type. using type assertion to check the type.
	var a interface{} = 1
	b := a.(int)
	fmt.Println(b) // 1
	//type assertion: for string.
	var c interface{} = "hello"
	d := c.(string)
	fmt.Println(d) // hello

	//if else, type assertion: if the type assertion fails, the program will panic. if the type assertion succeeds, the program will continue. in the below example: if e is not an int, the program will panic.
	var e interface{} = 1
	if i, ok := e.(int); ok {
		fmt.Println(i)
	}
}
