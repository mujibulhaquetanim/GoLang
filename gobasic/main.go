package main

import (
    "fmt"
    "github.com/mujibulhaquetanim/gobasic/struc"
    "github.com/mujibulhaquetanim/gobasic/datastructures"
)

func main() {
    p1 := struc.Person{Name: "Millie", Age: 20, Hobby: []string{"acting", "taking photos"}}
    fmt.Println(p1)

    datastructures.StringMethods()

}
