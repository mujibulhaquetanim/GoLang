package main

import (
	"encoding/json"
	"fmt"

	// "io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Welcome to CRUD using net/http")

	res, err := http.Get("http://jsonplaceholder.typicode.com/todos/5")
	if err != nil {
		fmt.Println("Error fetching data", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error response status code", res.StatusCode)
		return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
	fmt.Println("todo: ", todo)

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading bytes from response", err)
	// 	return
	// }

	// fmt.Println(string(data))

}
