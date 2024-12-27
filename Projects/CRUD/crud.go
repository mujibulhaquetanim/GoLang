package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	// "io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
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

func performPostRequest() {
	todo := Todo{
		UserId:    1,
		Id:        1,
		Title:     "CRUD using net/http",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling todo to JSON", err)
		return
	}

	// Convert byte array to string
	// jsonString := string(jsonData)

	// string.NewReader is used to convert string to io.Reader interface.
	// jsonReader := strings.NewReader(jsonString)
	jsonReader := bytes.NewReader(jsonData)

	res, err := http.Post("http://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error posting data", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		fmt.Println("Error response status code", res.StatusCode)
		return
	}
	fmt.Println("Data posted successfully", res.StatusCode)

	// to read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading bytes from response", err)
		return
	}
	fmt.Println(string(data))
}

func perforPutRequest() {
	updatedTodo := Todo{
		UserId:    2,
		Id:        2,
		Title:     "CRUD using net/http version 2 will be applied in another file",
		Completed: false,
	}

	jsonData, err := json.Marshal(updatedTodo)
	if err != nil {
		fmt.Println("Error marshalling update todo to JSON", err)
		return
	}

	// converting byte slice to io.Reader interface as http.NewRequest accepts io.Reader interface
	jsonReader := bytes.NewReader(jsonData)
	// jsonReader := strings.NewReader(string(jsonData))

	// 'Put' can be used instead of http.MethodPut
	req, err := http.NewRequest(http.MethodPut, "http://jsonplaceholder.typicode.com/todos/1", jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request", err)
		return
	}

	// setting the header for the request
	req.Header.Set("Content-Type", "application/json")

	// creating a client
	client := http.Client{}
	// sending the request and getting the response back
	res, err := client.Do((req))
	if err != nil {
		fmt.Println("Error updating data", err)
		return
	}
	defer res.Body.Close()

	resByteData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading bytes from response", err)
		return
	}

	fmt.Println("Response status code", res.StatusCode)
	fmt.Println("Data updated successfully", string(resByteData))
}

func performDeleteRequest() {
	req, err := http.NewRequest(http.MethodDelete, "http://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error creating Delete request", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting data", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error response status code", res.StatusCode)
		return
	}
	fmt.Println("Data deleted successfully", res.StatusCode)
}

func main() {
	fmt.Println("Welcome to CRUD using net/http")
	// performGetRequest()
	// performPostRequest()
	// perforPutRequest()
	performDeleteRequest()
}
