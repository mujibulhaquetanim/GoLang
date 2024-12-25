package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main(){
	instanceOfPersonStruct := Person{
		Name: "Emma Watson",
		Age: 31,
		IsAdult: true,
	}

	fmt.Println("Instance of Person struct:", instanceOfPersonStruct)

	jsonData, err := json.Marshal(instanceOfPersonStruct)
	if err != nil {
		fmt.Println("Error while marshalling:", err)
	}
	fmt.Println("JSON Data:", string(jsonData))

	// Unmarshalling JSON data to struct object Person which is convertedJsonData in this case.
	var convertedJsonData Person
	err = json.Unmarshal(jsonData, &convertedJsonData)
	if err != nil {
		fmt.Println("Error while unmarshalling:", err)
	}

	fmt.Println("Decoded JSON Data to struct object:", convertedJsonData)
}