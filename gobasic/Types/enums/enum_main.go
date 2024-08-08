package main

import "fmt"

//there is no type called enum in golang, but we can implement it using constants and types
//enum is a type that can have only one value
type enum int

const (
	zero enum = iota //iota is a counter that starts from 0 and increments by 1 for each constant in the enum type
	one
	two
)

func changeOrderStatus(status enum) {
	fmt.Println("order status: ", status)
}

func main() {
	changeOrderStatus(one) // 1
}
