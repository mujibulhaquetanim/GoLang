package main

import "fmt"

// CountDown iterator function takes in a value and returns a func that takes in another func with a signature of `func(int) bool`
func CountDown(v int) func(func(int) bool) {

	return func(yield func(int) bool) {
		for i := v; i > 0; i-- {
			// we then start a for loop that iterates
			if !yield(i) {
				// we then return and finish our iterations
				return
			}
		}
	}
}

//more complex example:
type Employee struct {
	name   string
	salary int
}

var employees = []Employee{
	{name: "John", salary: 1000},
	{name: "Jane", salary: 2000},
	{name: "Jim", salary: 3000},
}

func EmployeeIterator(e []Employee) func(func(int, Employee) bool) {
	return func(yield func(int, Employee) bool) {
		for i := 0; i < len(e); i++ {
			if !yield(i, e[i]) {
				return
			}
		}
	}
}

func main() {
	for x := range CountDown(5) {
		fmt.Println(x)
	}

	for x, e := range EmployeeIterator(employees) {
		fmt.Println(x, e)
	}
}
