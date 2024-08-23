package main

import (
	"fmt"
	"iter"
)

// iter.seq used instead of func(func(int) bool) syntax
func CountDown(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := v; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

// iter.seq2 example:
type Employee struct {
	name   string
	salary int
}

var employees = []Employee{
	{name: "John", salary: 1000},
	{name: "Jane", salary: 2000},
	{name: "Jim", salary: 3000},
}

func EmployeeIterator(e []Employee) iter.Seq2[int, Employee] {
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
