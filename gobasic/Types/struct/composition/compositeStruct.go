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
	// p1 := PersonInfo{} //by default the value of string is " ", int is 0, so the output will be: { 0 {  }}
	p1 := new(PersonInfo) //new returns a pointer, so we can use *p1 to access the value (dereferencing)
	p1.Name = "Millie bobby brown"
	p1.Age = 20
	//p1.Address = Address{"London", "England", "UK"}
	p1.Address = Address{
		City:    "London",
		State:   "England",
		Country: "UK",
	}
	fmt.Println(*p1)
	fmt.Println(p1.Address.City) //London //accessing the nested struct
	fmt.Print(p1.Address)
}
