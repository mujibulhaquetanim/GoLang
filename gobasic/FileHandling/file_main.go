package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter file name: ")

//take input from user and store in variable
	var fileName string
	fmt.Println("Enter file name: ")
	fmt.Scan(&fileName)

	//write it in a file and open it

	file, err := os.Create(fileName + ".txt")

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(file)

	defer file.Close()
}
