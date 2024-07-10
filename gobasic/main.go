package main

import (
    "fmt"
    "gobasic/struc" // Adjust relative path as per your GOPATH setup
)

func main() {
    p1 := struc.Person{Name: "Millie", Age: 20, Hobby: []string{"acting", "taking photos"}}

    fmt.Println(p1)
}
