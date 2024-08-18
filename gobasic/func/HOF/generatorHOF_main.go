package main

import "fmt"

func GenHOF(yield func(int) bool) {
	//infinite loop
	for i := 0; ; i++ {
		//calling the callback function passed as an argument and passing the value of i as an argument to it and checking if it returns false or not if it returns false it breaks the loop and returns from the function.
		if !yield(i) {
			return
		}
	}
}

func main() {
	//passing a callback function as an argument
	GenHOF(func(n int) bool {
		if n > 7 {
			fmt.Println("breaking the loop")
			return false
		}
		//printing the value of n
		fmt.Printf("yielding value in the loop from the callback: %d\n", n)
		return true
	})
}
