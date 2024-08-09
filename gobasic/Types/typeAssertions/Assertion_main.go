package main

import "fmt"

//interface
type Foer interface {
	Foo()
}
type MyStruct struct{}

func (f *MyStruct) Foo() {}
func (b *MyStruct) Bar() {}

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
	if onlyString, ok := e.(string); ok {
		fmt.Println("e is a string", onlyString)
	}
	var f Foer = &MyStruct{}
	if onlyMyStruct, ok := f.(interface{ Bar() }); ok {
		onlyMyStruct.Bar()
	}
}
