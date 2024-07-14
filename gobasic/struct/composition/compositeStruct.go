package main

import (
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
}

// composition: it is a struct that contains another struct. i.e. Address
type PersonInfo struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	// p1 := PersonInfo{"Millie bobby brown", 20, Address{"London", "England", "UK"}}
	p1 := PersonInfo{} //by default the value of string is " ", int is 0, so the output will be: { 0 {  }}

	p1.Name = "Millie bobby brown"
	p1.Age = 20
	p1.Address = Address{"London", "England", "UK"}
	fmt.Println(p1)
}
