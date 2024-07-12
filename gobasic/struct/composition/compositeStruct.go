//+build composition
package main

import (
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
}
//composition: it is a struct that contains another struct. i.e. Address
type PersonInfo struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	p1 := PersonInfo{"Millie bobby brown", 20, Address{"London", "England", "UK"}}
	fmt.Println(p1)
}
