package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to CRUD using net/http")

	res, err := http.Get("http://jsonplaceholder.typicode.com/todos/5")
	if err != nil {
		fmt.Println("Error fetching data", err)
		return
	}
	// Close the response body when done because it is a resource that must be released after use to prevent memory leaks. it not closed, it will be left open and the underlying connection will not be re-used by the http client for future requests. in golang it needs to close any resources that are opened to prevent memory leaks.
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error response status code", res.StatusCode)
		return
	}

	// ReadAll reads from the response body until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
	data, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading bytes from response", err)
		return
	}

	fmt.Println(string(data))

}
