package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	n := map[string]string{"name": "John", "surname": "Doe"}
	m["c"] = 3
	n["age"] = "30"
	fmt.Println(m)
	fmt.Println(n)
}
