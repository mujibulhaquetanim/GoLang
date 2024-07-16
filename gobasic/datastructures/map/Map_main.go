package main

import "fmt"

//generic map
//The comparable constraint ensures a type can be compared using == and !=.
// The K comparable constraint ensures that the type K (the key type) can be used in a map. Maps in Go require that their key types support equality comparison because keys need to be compared to determine if a key already exists in the map. Using comparable is practical for functions that need to ensure types can be compared for equality, such as when working with maps, sets, or other data structures that involve key lookups. This is required for key types in maps because the map operations (like checking if a key exists, inserting a new key, and deleting a key) rely on these comparison operators.
func PrintMap[K comparable, V any](m map[K]V) {
	for key, value := range m {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}

// KeyExists checks if a key exists in the map.
func KeyExists[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}

func main() {
	m := map[string]int{"a": 1, "b": 2}
	n := map[string]string{"name": "John", "surname": "Doe"}
	m["c"] = 3
	n["age"] = "30"
	fmt.Println(m)
	fmt.Println(n)

	intMap := map[string]int{
		"Emma":   34,
		"Millie": 20,
	}
	strMap := map[string]string{
		"name":    "Emma",
		"surname": "Watson",
	}
	floatMap := map[int]float64{
		1: 1.1,
		2: 2.2,
		3: 3.3,
	}

	PrintMap(intMap)
	PrintMap(strMap)
	PrintMap(floatMap)
	fmt.Println(KeyExists(strMap, "name")) //true
	fmt.Println(KeyExists(strMap, "job"))  //false
}
