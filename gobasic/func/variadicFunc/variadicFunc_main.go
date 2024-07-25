package main

import "fmt"

//while taking as parameters ... are used in the beginning of the type name, for taking argument ... used after the slice name (if slice used)
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("passing arguments directly: ",sum(1, 2, 3, 4, 5))

	nums:= []int{1,2,3,4,5}
	fmt.Println("passing arguments via slice: ", sum(nums...))
}
