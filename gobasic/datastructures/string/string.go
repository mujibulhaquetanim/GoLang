package main

import (
	"fmt"
	"strings"
)

func StringMethods() {
	name := "Mil lie"
	nameArr := strings.Split(name, " ")

	fmt.Println("Trimmed value is: ", strings.Trim("_Hello__world", "_"))

	fmt.Printf("index at 0: %s, index at 1: %s\n", nameArr[0], nameArr[1])
	fmt.Println(len(nameArr))
}

func main() {
	StringMethods()
}