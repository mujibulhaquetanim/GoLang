package main

import "fmt"

//generic map
//The comparable constraint ensures a type can be compared using == and !=.
// The K comparable constraint ensures that the type K (the key type) can be used in a map. Maps in Go require that their key types support equality comparison because keys need to be compared to determine if a key already exists in the map.
func PrintMap[K comparable, V any](m map[K]V) {
	for key, value := range m {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}

func main() {
	m := map[string]int{"a": 1, "b": 2}
	n := map[string]string{"name": "John", "surname": "Doe"}
	m["c"] = 3
	n["age"] = "30"
	fmt.Println(m)
	fmt.Println(n)

	intMapt := map[string]int{
		"Emma":   34,
		"Millie": 20,
	}
	strMap := map[string]string{
		"name":    "Emma",
		"surname": "Watson",
	}
	floatMap := map[string]float64{
		"height": 1.7,
		"weight": 60.5,
	}
	PrintMap(intMapt)
	PrintMap(strMap)
	PrintMap(floatMap)
}
