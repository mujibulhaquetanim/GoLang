package main

import "fmt"

//passing multiple return types
func funcky() (string, string, string, bool) {
	return "js", "go", "python", true
}

//receiving higher order function
func higherOrderFunc(fn func(high string) string) {
	fn("hi, from higher order function")
}

//returning function
func returnFunc() func(str string) string {
	return func(str string) string {
		return str
	}
}

func main() {
	//passing multiple return types
	lang1, lang2, lang3, bool := funcky()
	fmt.Printf("languages %t for : %s, %s, %s\n", bool, lang2, lang3, lang1)

	//passing higher order function (anonymous function)
	fn := func(str string) string {
		fmt.Println(str)
		return str
	}
	higherOrderFunc(fn)

	//returning function
	fn2 := returnFunc()
	fmt.Println(fn2("hello, from returning function"))
}
