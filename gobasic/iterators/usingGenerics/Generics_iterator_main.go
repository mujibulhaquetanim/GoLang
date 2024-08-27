package main

import (
	"fmt"
	"time"
	"iter"
)

type Employee struct {
	Name   string
	Salary int
}

// create a pre-defined list of employees  
var Employees = []Employee{
	{Name: "Elliot", Salary: 4},
	{Name: "Donna", Salary: 5},
}

	func Iterate[E any](e []E) func(func(int, E) bool) {
		return func(yield func(int, E) bool) {
			for i := 0; i <= len(e)-1; i++ {
				// to really hammer home the point about memory efficiency
				// let's add a sleep here to simulate us doing some fairly
				// intense computational work
				time.Sleep(5 * time.Second)
				if !yield(i, e[i]) {
					return
				}
			}
		}
	}
	func EmployeeIterate[E any](e []E) iter.Seq2[int, E] {
		return func(yield func(int, E) bool) {
			for i := 0; i <= len(e)-1; i++ {
				time.Sleep(5 * time.Second)
				if !yield(i, e[i]) {
					return
				}
			}
		}
	}

func main() {

	for x := range Iterate([]int{1, 2, 3}) {
		fmt.Println(x)
	}

	for x, e := range EmployeeIterate(Employees) {
		fmt.Println(x, e)
	}

}