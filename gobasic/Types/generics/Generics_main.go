package main

import "fmt"

func add[T int | float64](a, b T) T {
	return a + b
}

//comparable is a type that can be compared using ==, !=, <, >, <=, >= operators
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func min[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func addType[T float64](a T) T {
	return a
}

type subtype float64

//as the function only accepts type float64 but the given argument is of type subtype, which has underlying type float64, so we need to add ~ for any type as long as underlying type is float64
func addSubtype[T ~float64](a, b T) T {
	return a + b
}

func main() {
	printSlice([]int{1, 2, 3})
	printSlice([]string{"a", "b", "c"})

	var subType subtype = 1.5

	fmt.Println(add(1.5, 2))           // 3.5 //without type inference
	fmt.Println(add(1, 2))             // 3
	fmt.Println(addType[float64](1.5)) // 1.5 //with type inference

	fmt.Println(min(1.5, 2.5)) // 1.5

	fmt.Println(addSubtype(subType, 2.5)) // 4
}

// operator + not defined on a (variable of type T constrained by any
// func addType[T any] (a, b T) T {
// 	return a + b
// }
//any and interface are same
// func addSum[T interface{}](a,b T){

// 	return a + b
// }
