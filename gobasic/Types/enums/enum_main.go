package main

import "fmt"

//there is no type called enum in golang, but we can implement it using constants and types
//enum is a type that can have only one value
type OrderStatus int

const (
	Received OrderStatus = iota //iota is a counter that starts from 0 and increments by 1 for each constant in the enum type
	Confirmed
	Prepared
	Dispatched
	Shipped
)

type ProductType string

const (
	Book     ProductType = "book"
	Toy      ProductType = "toy"
	Pen                  = "pen" //we can write without the type
	Notebook             = "notebook"
)

func printProductType(productType ProductType) {
	fmt.Println("product type: ", productType)
}

func changeOrderStatus(status OrderStatus) {
	fmt.Println("order status: ", status)
}

func main() {
	changeOrderStatus(Confirmed) // 1

	printProductType(Pen) // pen

}
