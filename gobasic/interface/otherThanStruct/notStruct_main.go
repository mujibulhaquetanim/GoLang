package main

import "fmt"

type Counter interface {
	Count()
}

type Number int
type String string

func (n Number) Count() int {
	counter := 0
	if n > 0 {
		counter++
	}
	return counter
}

func (s String) Count() int {
	return len(s)
}

func main() {
	w := String("hello")
	n := Number(5)
	fmt.Println(w.Count())
	fmt.Println(n.Count())
}
